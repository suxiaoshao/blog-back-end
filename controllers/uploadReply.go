package controllers

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/model"
	"nextBlogServer/util"
	"regexp"
)

func UploadReply(context *gin.Context) {
	//绑定上传数据
	query := struct {
		Aid     int64  `form:"aid"`
		Content string `form:"content"`
		Email   string `form:"email"`
		Name    string `form:"name"`
		Url     string `form:"url"`
	}{
		Aid:     0,
		Content: "",
		Email:   "",
		Name:    "",
		Url:     "",
	}
	err := context.BindJSON(&query)

	//验证 email 是否合法
	emailRegex, _ := regexp.Compile("^([A-Za-z0-9_\\-.])+@([A-Za-z0-9_\\-.])+\\.([A-Za-z]{2,4})$")

	//判断 url 是否合法
	urlRegex, _ := regexp.Compile("(^$|^(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]$)")
	if err != nil || query.Aid == 0 || query.Content == "" || query.Name == "" ||
		!emailRegex.MatchString(query.Email) || !urlRegex.MatchString(query.Url) {
		context.JSON(200, util.ReturnMessageError("上传数据错误", gin.H{}))
		return
	}
	// 获取文章
	article, err := model.GetArticleContentByAid(query.Aid)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("文章不存在", gin.H{}))
		return
	}
	// 获取新评论
	newReply, err := article.AddNewReply(query.Content, query.Name, query.Email, query.Url)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据保存错误", gin.H{}))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(newReply))
}
