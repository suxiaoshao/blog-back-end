package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
)

func ArticleNum(context *gin.Context) {
	query := struct {
		SearchName string  `form:"searchName"`
		Type       []int64 `form:"type"`
		AllType    bool    `form:"allType"`
	}{SearchName: "", Type: []int64{}, AllType: false}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("query数据错误", gin.H{"count": 0}))
		return
	}

	searchQuery := model.GetSearchQuery(0, 0, query.SearchName, query.Type, -1, query.AllType)
	count, err := searchQuery.GetArticleNum()
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", gin.H{"count": 0}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"count": count}))
}
