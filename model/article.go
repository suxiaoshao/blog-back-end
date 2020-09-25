package model

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ArticleItem struct {
	Title     string  `bson:"title" json:"title"`
	Aid       int64   `bson:"aid" json:"aid"`
	Type      []int64 `bson:"type" json:"type"`
	ReplyNum  int64   `bson:"reply_num" json:"replyNum"`
	ReadNum   int64   `bson:"read_num" json:"readNum"`
	TimeStamp int64   `bson:"time_stamp" json:"timeStamp"`
}

type ArticleContent struct {
	Title     string  `bson:"title" json:"title"`
	Aid       int64   `bson:"aid" json:"aid"`
	Type      []int64 `bson:"type" json:"type"`
	ReplyNum  int64   `bson:"reply_num" json:"replyNum"`
	ReadNum   int64   `bson:"read_num" json:"readNum"`
	Content   string  `bson:"content" json:"content"`
	TimeStamp int64   `bson:"time_stamp" json:"timeStamp"`
}

// 阅读数加一
func (article *ArticleContent) Read() *ArticleContent {
	article.ReadNum = article.ReadNum + 1
	_, _ = articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": article.Aid}, gin.H{"$set": gin.H{"read_num": article.ReadNum}})
	return article
}

// 更新
func (article *ArticleContent) Update(content string, typeList []int64, title string) (*ArticleContent, error) {
	//获取时间字符串和时间
	timeStamp := time.Now().UnixNano() / 1000000
	newArticle := getArticleContent(title, article.Aid, typeList, content, timeStamp, article.ReplyNum, article.ReadNum)
	_, err := articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": article.Aid}, gin.H{"$set": newArticle})
	if err != nil {
		return nil, err
	}
	return newArticle, nil
}

// 写入数据库
func (article *ArticleContent) WriteToDatabase() (*ArticleContent, error) {
	_, err := articleDatabase.InsertOne(context.TODO(), article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

//添加新评论
func (article *ArticleContent) AddNewReply(content string, name string, email string, url string) (*ReplyItem, error) {
	newReply, err := getNewReply(article.Aid, content, name, email, url)
	if err != nil {
		return nil, err
	}
	newReply, err = newReply.writeToDatabase()
	if err != nil {
		return nil, err
	}
	return newReply, nil
}

// 获取评论链接
func (article *ArticleContent) GetReplyList(offset int64, limit int64) ([]ReplyItem, error) {
	//限制limit和offset的范围
	if limit > 50 || limit < 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	Result := ReplyItem{
		Aid:     0,
		Content: "",
		Email:   "",
		Name:    "",
		Rid:     -1,
		Url:     "",
	}
	resultList := make([]ReplyItem, 0)

	//生成 findOption
	findOption := options.Find()
	findOption.SetSkip(offset)
	findOption.SetLimit(limit)

	cur, err := replyDatabase.Find(context.TODO(), gin.H{"aid": article.Aid}, findOption)
	if err != nil {
		return resultList, err
	}
	for cur.Next(context.TODO()) {
		result := Result
		err = cur.Decode(&result)
		if err == nil && result.Rid != -1 {
			resultList = append(resultList, result)
		}
	}
	return resultList, nil
}

// 获取评论数
func (article *ArticleContent) GetReplyNum() (int64, error) {
	count, err := replyDatabase.CountDocuments(context.TODO(), gin.H{"aid": article.Aid})
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 获取老
func GetArticleContentByAid(aid int64) (*ArticleContent, error) {
	article := getArticleContent("", 0, []int64{}, "", 0, 0, 0)
	err := articleDatabase.FindOne(context.TODO(), gin.H{"aid": aid}).Decode(article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// 获取
func getArticleContent(title string, aid int64, typeList []int64, content string, timeStamp int64, replyNum int64, readNum int64) *ArticleContent {
	return &ArticleContent{
		Title:     title,
		Aid:       aid,
		Type:      typeList,
		ReplyNum:  replyNum,
		ReadNum:   readNum,
		Content:   content,
		TimeStamp: timeStamp,
	}
}

// 创建新文章
func CreateNewArticle(title string, typeList []int64, content string) (*ArticleContent, error) {
	//获取时间字符串和时间
	timeStamp := time.Now().UnixNano() / 1000000
	count, err := articleDatabase.CountDocuments(context.TODO(), gin.H{})
	if err != nil {
		return nil, err
	}
	return getArticleContent(title, count+1, typeList, content, timeStamp, 0, 0), nil
}
