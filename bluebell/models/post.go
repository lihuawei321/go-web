package models

import (
	"time"
)

type Post struct {
	ID          int64     `json:"post_id,string" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

//func (p *Post) UnmarshalJSON(data []byte) (err error) {
//	required := struct {
//		Title       string `json:"title" db:"title"`
//		Content     string `json:"content" db:"content"`
//		CommunityID int64  `json:"community_id" db:"community_id"`
//	}{}
//	err = json.Unmarshal(data, &required)
//	if err != nil {
//		return
//	} else if len(required.Title) == 0 {
//		err = errors.New("帖子标题不能为空")
//	} else if len(required.Content) == 0 {
//		err = errors.New("帖子内容不能为空")
//	} else if required.CommunityID == 0 {
//		err = errors.New("未指定版块")
//	} else {
//		p.Title = required.Title
//		p.Content = required.Content
//		//p.CommunityID = required.CommunityID
//	}
//	return
//}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	VoteNum          int64                     `json:"vote_num"`
	*Post                                      //嵌入帖子结构体
	*CommunityDetail `json:"community_detail"` //嵌入社区结构体
}
