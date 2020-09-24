package model

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type ReplyItem struct {
	Aid       int64  `bson:"aid" json:"aid"`
	Content   string `bson:"content" json:"content"`
	Email     string `bson:"email" json:"email"`
	Name      string `bson:"name" json:"name"`
	Rid       int64  `bson:"rid" json:"rid"`
	Url       string `bson:"url" json:"url"`
	TimeStamp int64  `bson:"time_stamp" json:"timeStamp"`
}

// 获取新评论
func getNewReply(aid int64, content string, name string, email string, url string) (*ReplyItem, error) {
	//获取时间字符串和时间
	timeStamp := time.Now().UnixNano() / 1000000
	//更新 article 的 reply_num 字段
	updateResult, err := articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": aid}, gin.H{"$inc": gin.H{"reply_num": 1}})
	if err != nil || updateResult.MatchedCount == 0 {
		return nil, errors.New("文章不存在")
	}
	// 获取评论数
	rid, err := replyDatabase.CountDocuments(context.TODO(), gin.H{})
	if err != nil {
		return nil, err
	}
	rid++
	return &ReplyItem{
		Aid:       aid,
		Content:   content,
		Email:     email,
		Name:      name,
		Rid:       rid,
		Url:       url,
		TimeStamp: timeStamp,
	}, nil
}
// 评论写入
func (reply *ReplyItem) writeToDatabase() (*ReplyItem, error) {
	_, err := replyDatabase.InsertOne(context.TODO(), &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}


