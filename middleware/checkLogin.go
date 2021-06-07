package middleware

import (
	"blogServer/model"
	"blogServer/util"
	"github.com/gin-gonic/gin"
)

// CheckLogin /* 判断是否登入 */
func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		ip := context.ClientIP()
		loginClaims, err := util.Decryption(token, ip)
		if err != nil {
			context.Abort()
			context.JSON(301, util.ReturnMessageError("未登录", nil))
			return
		}
		_, err = model.UserManager.GetUserWithCheck(loginClaims.UserName, loginClaims.Password)
		if err != nil {
			context.Abort()
			context.JSON(301, util.ReturnMessageError("登录过期", nil))
			return
		}
		context.Next()
	}
}
