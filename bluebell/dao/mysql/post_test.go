package mysql

import (
	"go-web/bluebell/models"
	"go-web/bluebell/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "12345678",
		DB:           "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          12322,
		AuthorId:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just test content",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed: %v", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
