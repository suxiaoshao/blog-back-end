package controllers

import (
	"blogServer/model"
	"blogServer/util"
	"github.com/gin-gonic/gin"
)

type ArticleListQuery struct {
	LabelIds     []uint `form:"labelIds"`
	PreArticleId uint   `form:"preArticleId"`
	Limit        uint   `form:"limit"`
}

func ArticleList(context *gin.Context) {
	query := ArticleListQuery{
		LabelIds:     []uint{},
		Limit:        20,
		PreArticleId: 0,
	}
	// 解析上传内容
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("上传数据错误", gin.H{"count": 0}))
		return
	}
	// 获取数据
	articleInfos, err := model.ArticleManager.GetArticleInfoList(query.LabelIds, query.Limit, query.PreArticleId)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", nil))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"articleList": articleInfos}))
}
