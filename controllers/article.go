package controllers

import (
	"blogServer/model"
	"blogServer/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticleDetail 获取文章具体信息
func GetArticleDetail(context *gin.Context) {
	params := context.Param("articleId")
	// 解析articleId
	aid, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		context.JSON(400, util.ReturnMessageError("上传数据错误", gin.H{}))
		return
	}
	// 获取详细数据
	article, err := model.ArticleManager.GetArticleByArticleId(uint(aid))
	if err != nil {
		context.JSON(404, util.ReturnMessageError("没有此数据", gin.H{}))
		return
	}
	articleDetail, err := article.GetArticleDetail()
	if err != nil {
		context.JSON(500, util.ReturnMessageError("服务器内部错误", gin.H{}))
		return
	}
	// 添加阅读记录
	ip := context.ClientIP()
	_ = article.ReadAdd(ip)
	context.JSON(200, util.ReturnMessageSuccess(articleDetail))
}

// 上传数据
type updateArticle struct {
	Content  string `form:"content"`
	Title    string `form:"title"`
	LabelIds []uint `form:"labelIds"`
	Aid      int64  `form:"aid"`
}

// UpdateArticle 更新文章信息
func UpdateArticle(context *gin.Context) {
	// 验证参数
	query := updateArticle{Content: "", Title: "", LabelIds: make([]uint, 0), Aid: 0}
	err := context.BindJSON(&query)
	if err != nil || query.Content == "" || query.Title == "" || model.LabelManager.HasLabels(query.LabelIds) {
		context.JSON(200, util.ReturnMessageError("上传数据错误", nil))
		return
	}
	// 如果文章还不存在
	if query.Aid == 0 {
		err := model.ArticleManager.AddArticle(query.Title, query.Content, query.LabelIds)
		if err != nil {
			context.JSON(500, util.ReturnMessageError("添加文章失败",nil))
			return
		}
		context.JSON(200, util.ReturnMessageSuccess(nil))
		return
	}
	// 修改文章
	err = model.ArticleManager.UpdateArticle(uint(query.Aid), query.Title, query.Content, query.LabelIds)
	if err != nil {
		context.JSON(200, util.ReturnMessageError("修改文章失败", nil))
		return
	}
	context.JSON(200, util.ReturnMessageSuccess(nil))
}
