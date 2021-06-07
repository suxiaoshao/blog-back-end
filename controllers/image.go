package controllers

import (
	"blogServer/util"
	"github.com/gin-gonic/gin"
)

func ImageUpload(context *gin.Context) {
	f, err := context.FormFile("image")
	if err != nil {
		context.JSON(200, util.ReturnMessageError("上传数据错误", gin.H{}))
		return
	}
	file, err := f.Open()
	if err != nil {
		context.JSON(200, util.ReturnMessageError("上传数据错误", gin.H{}))
		return
	}
	buf := make([]byte, f.Size)
	_, _ = file.Read(buf)
	_, _ = context.Writer.WriteString(string(buf))
}
