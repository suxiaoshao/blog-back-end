package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
)

func Wallpaper (context *gin.Context) {
	imageString, err := model.GetRandomImage()
	if err != nil {
		context.JSON(404, gin.H{})
		return
	}
	_, _ = context.Writer.WriteString(imageString)
}
