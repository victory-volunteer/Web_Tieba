package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

func init() { //初始化mysql函数（防止空指针db引用）
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "root123",
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

func TestCreatePost(t *testing.T) { //测试能否正常的将数据插到数据库中
	post := models.Post{
		ID:          12,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err) //Fatalf会自动将测试用例报错并停下来
	}
	t.Logf("CreatePost insert record into mysql success") //Logf用来输出
}
