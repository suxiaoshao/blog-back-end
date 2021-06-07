package model

import (
	"blogServer/database"
)

var ArticleReadManger = ArticleReadDao{}

// ArticleRead 访问数据的数据库
type ArticleRead struct {
	ReadId    uint   `gorm:"read_id;primaryKey;not null"`
	Ip        string `gorm:"ip;not null"`
	ArticleId uint   `gorm:"article_id;not null"`
}

func (article ArticleRead) TableName() string {
	return "article_read"
}

type ArticleReadDao struct{}

// GetReadNumByArticleId 根据 articleId 判断阅读数
func (articleReadDao ArticleReadDao) GetReadNumByArticleId(articleId uint) (*int64, error) {
	count := new(int64)
	result := database.MysqlDb.Model(ArticleRead{}).Where("article_id = ?", articleId).Count(count)
	if result.Error != nil {
		return nil, result.Error
	}
	return count, nil
}

// AddArticleRead 添加一个阅读记录
func (articleReadDao ArticleReadDao) AddArticleRead(ip string, articleId uint) error {
	articleRead := ArticleRead{ArticleId: articleId, Ip: ip}
	result := database.MysqlDb.Create(&articleRead)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
