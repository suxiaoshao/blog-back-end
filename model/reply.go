package model

import (
	"blogServer/database"
	"blogServer/util"
	"errors"
	"time"
)

var ReplyManager = ReplyDao{}

type Reply struct {
	ArticleId  uint      `gorm:"article_id;not null"`
	Content    string    `gorm:"content;not null"`
	ReplyId    uint      `gorm:"reply_id;not null"`
	Email      string    `gorm:"email;not null"`
	Name       string    `gorm:"name;not null"`
	CreateTime time.Time `gorm:"create_time;not null"`
	url        *string   `gorm:"url"`
	Ip         string    `gorm:"ip;not null"`
}

func (reply Reply) TableName() string {
	return "reply"
}

type ReplyDao struct{}

// GetReplyByArticleId 获取文章评论
func (replyDao ReplyDao) GetReplyByArticleId(articleId uint, preReplyId uint, limit int) ([]Reply, error) {
	limit = util.If(limit > 50, 50, limit).(int)
	var replies []Reply
	result := database.MysqlDb.Where("article_id = ? AND reply_id > ?", articleId, preReplyId).Order("reply_id DESC").Limit(limit).Find(&replies)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(replies) == 0 {
		return nil, errors.New("没有更多数据了")
	}
	return replies, nil
}

// GetReplyNumByArticleId 获取文章评论数
func (replyDao ReplyDao) GetReplyNumByArticleId(articleId uint) (*int64, error) {
	count := new(int64)
	result := database.MysqlDb.Model(Reply{}).Where("article_id = ?", articleId).Count(count)
	if result.Error != nil {
		return nil, result.Error
	}
	return count, nil
}

// AddNewReply 新建评论
func (replyDao ReplyDao) AddNewReply(articleId uint, content string, email string, name string, ip string, url *string) (*Reply, error) {
	// 添加 reply 数据
	currentTime := time.Now()
	reply := Reply{
		ArticleId:  articleId,
		Content:    content,
		ReplyId:    0,
		Email:      email,
		Name:       name,
		CreateTime: currentTime,
		url:        url,
		Ip:         ip,
	}
	tx := database.MysqlDb.Create(&reply)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &reply, nil
}
