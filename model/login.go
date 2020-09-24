package model

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
)

func Login(user string, password string) (*LoginInfo, error) {
	info := LoginInfo{
		Password: "",
		User:     "",
	}
	infoDatabase := Client.Database("my_blog").Collection("info")
	err := infoDatabase.FindOne(context.TODO(), gin.H{}).Decode(&info)
	if err != nil || info.Password != password || info.User != user {
		return nil, errors.New("密码错误")
	}
	return &info, nil
}
