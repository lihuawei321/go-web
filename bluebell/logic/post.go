package logic

import (
	"go-web/bluebell/dao/mysql"
	"go-web/bluebell/models"
	"go-web/bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	return mysql.CreatePost(p)
	// 3. 返回
}
