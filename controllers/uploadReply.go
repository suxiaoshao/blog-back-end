package controllers

import (
	"blogServer/model"
	"blogServer/util"
	"github.com/gin-gonic/gin"
	"regexp"
)

type UploadReplyQuery struct {
	ArticleId int64   `form:"articleId" json:"articleId"`
	Content   string  `form:"content" json:"content"`
	Email     string  `form:"email" json:"email"`
	Name      string  `form:"name" json:"name"`
	Url       *string `form:"url" json:"url"`
}

func UploadReply(context *gin.Context) {
	//绑定上传数据
	query := UploadReplyQuery{
		ArticleId: 0,
		Content:   "",
		Email:     "",
		Name:      "",
		Url:       nil,
	}
	err := context.BindJSON(&query)

	//验证 email 是否合法
	emailRegex, _ := regexp.Compile("^([A-Za-z0-9_\\-.])+@([A-Za-z0-9_\\-.])+\\.([A-Za-z]{2,4})$")

	//判断 url 是否合法
	urlRegex, _ := regexp.Compile("(^$|^(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]$)")
	if err != nil || query.ArticleId == 0 || query.Content == "" || query.Name == "" ||
		!emailRegex.MatchString(query.Email) || !urlRegex.MatchString(*query.Url) {
		context.JSON(200, util.ReturnMessageError("上传数据错误", nil))
		return
	}
	// 获取新评论
	newReply, err := model.ReplyManager.AddNewReply(uint(query.ArticleId), query.Content, query.Email, query.Name, context.ClientIP(), query.Url)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("数据保存错误", nil))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(newReply))
}
