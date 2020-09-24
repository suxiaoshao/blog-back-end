package util

import "github.com/gin-gonic/gin"

func ReturnMessageError(msg string, data interface{}) gin.H {
	return gin.H{
		"success": false,
		"msg":     msg,
		"data":    data,
	}
}
func ReturnMessageSuccess(data interface{}) gin.H {
	return gin.H{
		"success": true,
		"msg":     "",
		"data":    data,
	}
}
