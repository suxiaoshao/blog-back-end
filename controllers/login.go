package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/middleware"
	"nextBlogServer/model"
	"nextBlogServer/util"
)

func Login(context *gin.Context) {
	/* 获取用户 ip 地址 */
	userIp := context.ClientIP()
	query := model.LoginInfo{
		Password: "",
		User:     "",
	}
	err := context.BindJSON(&query)
	if err != nil || query.Password == "" || query.User == "" {
		middleware.ClearCookie(context)
		context.JSON(200, util.ReturnMessageError("上传数据缺失", gin.H{}))
		return
	}
	loginInfo, err := model.Login(query.User, query.Password)
	if err != nil {
		middleware.ClearCookie(context)
		context.JSON(200, util.ReturnMessageError("账号密码错误", gin.H{}))
		return
	}
	uid, err := util.Encryption(loginInfo.User, userIp)
	if err != nil {
		middleware.ClearCookie(context)
		context.JSON(200, util.ReturnMessageError("服务器错误", gin.H{}))
		return
	}
	context.SetCookie("uid", uid, 0, "/api/admin", context.Request.Host, false, false)
	pid, err := util.Encryption(loginInfo.Password, userIp)
	if err != nil {
		middleware.ClearCookie(context)
		context.JSON(200, util.ReturnMessageError("服务器错误", gin.H{}))
		return
	}
	context.SetCookie("pid", pid, 0, "/api/admin", context.Request.Host, false, false)
	context.JSON(200, util.ReturnMessageSuccess(gin.H{}))
}
