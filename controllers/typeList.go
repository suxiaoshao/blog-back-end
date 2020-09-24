package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
)

func TypeList (context *gin.Context) {
	typeList, err := model.GetTypeList()
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取错误", gin.H{"typeList": make([]string, 0)}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"typeList": typeList}))
}
