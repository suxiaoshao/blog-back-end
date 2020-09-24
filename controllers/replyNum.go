package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
	"strconv"
)

func RePlyNum(context *gin.Context) {
	params := context.Param("aid")
	aid, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("url链接错误", gin.H{"count": 0}))
		return
	}

	// 获取文章
	article, err := model.GetArticleContentByAid(aid)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("文章不存在", gin.H{"count": 0}))
		return
	}
	// 获取文章评论数
	num, err := article.GetReplyNum()
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", gin.H{"count": 0}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"count": num}))
}
