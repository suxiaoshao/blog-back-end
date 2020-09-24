package middleware

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
)

/* 判断是否登入 */
func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		isNotLoginResponse := util.ReturnMessageError("未登入", gin.H{})
		uid, err := context.Cookie("uid")
		if err != nil {
			ClearCookie(context)
			context.Abort()
			context.JSON(200, isNotLoginResponse)
			return
		}
		/* 获取用户 ip 地址 */
		userIp := context.ClientIP()
		realUid, err := util.Decryption(uid, userIp)
		if err != nil {
			ClearCookie(context)
			context.Abort()
			context.JSON(200, isNotLoginResponse)
			return
		}
		pid, err := context.Cookie("pid")
		if err != nil {
			ClearCookie(context)
			context.Abort()
			context.JSON(200, isNotLoginResponse)
			return
		}
		realPid, err := util.Decryption(pid, userIp)
		if err != nil {
			ClearCookie(context)
			context.Abort()
			context.JSON(200, isNotLoginResponse)
			return
		}
		_, err = model.Login(realUid, realPid)
		if err != nil {
			ClearCookie(context)
			context.Abort()
			context.JSON(200, isNotLoginResponse)
			return
		}
		context.Next()
	}
}
func ClearCookie(context *gin.Context) {
	context.SetCookie("uid", "", -1, "/api/admin", context.Request.Host, false, false)
	context.SetCookie("pid", "", -1, "/api/admin", context.Request.Host, false, false)
}
