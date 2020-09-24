package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
)

func ArticleList(context *gin.Context) {
	query := struct {
		SearchName string  `form:"searchName"`
		Type       []int64 `form:"type"`
		Offset     int64   `form:"offset"`
		Limit      int64   `form:"limit"`
		Sort       int64   `form:"sort"`
		AllType    bool    `form:"allType"`
	}{
		SearchName: "",
		Type:       []int64{},
		Offset:     0,
		Limit:      20,
		Sort:       -1,
		AllType:    false,
	}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("query数据错误", gin.H{"count": 0}))
		return
	}

	searchQuery:=model.GetSearchQuery(query.Offset, query.Limit, query.SearchName, query.Type, query.Sort, query.AllType)
	articleList, err := searchQuery.GetArticleList()
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据获取失败", gin.H{"articleList": make([]model.ArticleItem, 0)}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(gin.H{"articleList": articleList}))
}
