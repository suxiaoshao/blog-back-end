package model

import "blogServer/database"

var ArticleLabelManager = ArticleLabelDao{}

type ArticleLabel struct {
	LabelId   uint `gorm:"label_id;primaryKey;autoIncrement:false;not null"`
	ArticleId uint `gorm:"article_id;primaryKey;autoIncrement:false;not null"`
}

func (article ArticleLabel) TableName() string {
	return "article_label"
}

type ArticleLabelDao struct{}

// ClearByArticleId 删除所有 articleId 的 label
func (articleLabelDao ArticleLabelDao) ClearByArticleId(articleId uint) error {
	db := database.MysqlDb.Where("article_id = ?", articleId).Delete(ArticleLabel{})
	return db.Error
}

// AddLabel 添加 label
func (articleLabelDao ArticleLabelDao) AddLabel(articleId uint, labelId uint) error {
	db := database.MysqlDb.Create(ArticleLabel{
		LabelId:   labelId,
		ArticleId: articleId,
	})
	return db.Error
}
