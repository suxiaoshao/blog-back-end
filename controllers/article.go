package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
	"strconv"
)

func GetArticle(context *gin.Context) {
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
func UpdateArticle(context *gin.Context) {
	query := struct {
		Content  string  `form:"content"`
		Title    string  `form:"title"`
		TypeList []int64 `form:"typeList"`
		Aid      int64   `form:"aid"`
	}{Content: "", Title: "", TypeList: make([]int64, 0), Aid: 0}
	err := context.BindJSON(&query)
	if err != nil || query.Content == "" || query.Title == "" || len(query.TypeList) == 0 {
		context.JSON(200, util.ReturnMessageError("上传数据错误", gin.H{}))
		return
	}
	if query.Aid == 0 {
		article, err := model.CreateNewArticle(query.Title, query.TypeList, query.Content)
		if err != nil {
			context.JSON(200, util.ReturnMessageError("数据库错误", gin.H{}))
			return
		}
		article, err = article.WriteToDatabase()
		if err != nil {
			context.JSON(200, util.ReturnMessageError("数据库错误", gin.H{}))
			return
		}
		context.JSON(200, util.ReturnMessageSuccess(article))
		return
	}
	article, err := model.GetArticleContentByAid(query.Aid)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("无法找到文章", gin.H{}))
		return
	}
	article, err = article.Update(query.Content, query.TypeList, query.Title)
	if err != nil {
		fmt.Println(11111)
		context.JSON(200, util.ReturnMessageError("数据库错误", gin.H{}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(article))
}
