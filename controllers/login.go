package controllers

import (
	"blogServer/model"
	"blogServer/util"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	// 验证上传数据
	userIp := context.ClientIP()
	query := model.User{}
	err := context.BindJSON(&query)
	if err != nil || query.Password == "" || query.UserName == "" {
		context.JSON(300, util.ReturnMessageError("上传数据缺失", nil))
		return
	}
	// 验证账号密码
	user, err := model.UserManager.GetUserWithCheck(query.UserName, query.Password)
	if err != nil {
		context.JSON(301, util.ReturnMessageError("账号密码错误", gin.H{}))
		return
	}
	token, err := util.Encryption(user.UserName, user.Password, userIp)
	if err != nil {
		context.JSON(500, util.ReturnMessageError("服务器内部错误", gin.H{}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"token": token}))
}
