package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
	"strconv"
)

func Article(context *gin.Context) {
	params := context.Param("aid")
	aid, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("params错误", gin.H{}))
		return
	}
	articleData, err := model.GetArticleContentByAid(aid)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", gin.H{}))
		return
	}
	articleData = articleData.Read()
	context.JSON(200, util.ReturnMessageSuccess(articleData))
}
