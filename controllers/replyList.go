package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
	"strconv"
)

func ReplyList(context *gin.Context) {
	//数据aid
	params := context.Param("aid")
	aid, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("url链接错误", gin.H{"replyList": make([]model.ReplyItem, 0)}))
		return
	}

	//offset limit数据获取
	query := struct {
		Offset int64 `form:"offset"`
		Limit  int64 `form:"limit"`
	}{
		Offset: 0,
		Limit:  20,
	}
	err = context.BindQuery(&query)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("query数据错误", gin.H{"replyList": make([]model.ReplyItem, 0)}))
		return
	}
	// 获取文章
	article,err:=model.GetArticleContentByAid(aid)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("文章不存在", gin.H{"replyList": make([]model.ReplyItem, 0)}))
		return
	}
	// 获取文章列表
	resultList, err := article.GetReplyList(query.Offset, query.Limit)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", gin.H{"replyList": make([]model.ReplyItem, 0)}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"replyList": resultList}))
}
