package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Post struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorId    uint64    `json:"author_id" db:"author_id" binding:"required"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title" db:"title"`
		Content     string `json:"content" db:"content"`
		CommunityID int64  `json:"community_id" db:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == 0 {
		err = errors.New("未指定版块")
	} else {
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityID = required.CommunityID
	}
	return
}

type ApiPostDetail struct {
	*Post
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}
