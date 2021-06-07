package controllers

import (
	"blogServer/model"
	"blogServer/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ReplyListQuery struct {
	PreReplyId uint `form:"preReplyId"`
	Limit      int  `form:"limit"`
}

func ReplyList(context *gin.Context) {
	//数据aid
	params := context.Param("articleId")
	articleId, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("上传数据错误", nil))
		return
	}

	//offset limit数据获取
	query := ReplyListQuery{
		PreReplyId: 0,
		Limit:      20,
	}
	err = context.BindQuery(&query)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("上传数据错误", nil))
		return
	}

	// 获取文章列表
	replyList, err := model.ReplyManager.GetReplyByArticleId(uint(articleId), query.PreReplyId, query.Limit)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", nil))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"replyList": replyList}))
}
