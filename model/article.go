package model

import (
	"blogServer/database"
	"blogServer/util"
	"errors"
	"time"
)

var ArticleManager = ArticleDao{}

// ArticleDetail 文章详细信息
type ArticleDetail struct {
	ArticleId  uint      `json:"articleId"`
	Content    string    `json:"content"`
	Title      string    `json:"title"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	ReadNum    int64     `json:"readNum"`
	ReplyNum   int64     `json:"replyNum"`
	Labels     []Label   `json:"labels"`
}

// ArticleInfo 文章列表信息
type ArticleInfo struct {
	ArticleId  uint      `json:"articleId"`
	Title      string    `json:"title"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	ReadNum    int64     `json:"readNum"`
	ReplyNum   int64     `json:"replyNum"`
	Labels     []Label   `json:"labels"`
}

// Article 数据库中的文章数据
type Article struct {
	ArticleId  uint      `gorm:"article_id;primaryKey;not null"`
	Content    string    `gorm:"content;not null"`
	Title      string    `gorm:"title;not null"`
	CreateTime time.Time `gorm:"create_time;not null"`
	UpdateTime time.Time `gorm:"update_time;not null"`
}

// TableName 绑定表名
func (article Article) TableName() string {
	return "article"
}

// GetArticleDetail 获取文章详细信息
func (article Article) GetArticleDetail() (*ArticleDetail, error) {
	articleInfo, err := article.GetArticleInfo()
	if err != nil {
		return nil, err
	}
	return &ArticleDetail{
		ArticleId:  articleInfo.ArticleId,
		Title:      articleInfo.Title,
		CreateTime: articleInfo.CreateTime,
		UpdateTime: articleInfo.UpdateTime,
		ReplyNum:   articleInfo.ReplyNum,
		ReadNum:    articleInfo.ReadNum,
		Labels:     articleInfo.Labels,
		Content:    article.Content,
	}, nil
}

// GetArticleInfo 获取文章详细信息
func (article Article) GetArticleInfo() (*ArticleInfo, error) {
	readNum, err := ArticleReadManger.GetReadNumByArticleId(article.ArticleId)
	if err != nil {
		return nil, err
	}
	replyNum, err := ReplyManager.GetReplyNumByArticleId(article.ArticleId)
	if err != nil {
		return nil, err
	}
	labels, err := LabelManager.GetLabelsByArticleId(article.ArticleId)
	if err != nil {
		return nil, err
	}
	return &ArticleInfo{
		ArticleId:  article.ArticleId,
		Title:      article.Title,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
		ReplyNum:   *replyNum,
		ReadNum:    *readNum,
		Labels:     labels,
	}, nil
}

// ReadAdd 阅读数加一
func (article Article) ReadAdd(ip string) error {
	err := ArticleReadManger.AddArticleRead(ip, article.ArticleId)
	return err
}

// ArticleDao 文章操作
type ArticleDao struct {
}

// GetArticleByArticleId 通过 id 获取文章数据
func (articleDao ArticleDao) GetArticleByArticleId(articleId uint) (*Article, error) {
	var article = new(Article)
	result := database.MysqlDb.First(article, articleId)
	if result.Error != nil {
		return nil, result.Error
	}
	return article, nil
}

// AddArticle 添加 article 数据
func (articleDao ArticleDao) AddArticle(title string, content string, labelIds []uint) error {
	// 添加 article 数据
	currentTime := time.Now()
	article := Article{
		Content:    content,
		Title:      title,
		CreateTime: currentTime,
		UpdateTime: currentTime,
	}
	db := database.MysqlDb.Model(article).Create(&article)
	if db.Error != nil {
		return db.Error
	}
	// 添加标签
	for _, labelId := range labelIds {
		err := ArticleLabelManager.AddLabel(article.ArticleId, labelId)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateArticle 更新 article 数据
func (articleDao ArticleDao) UpdateArticle(articleId uint, title string, content string, labelIds []uint) error {
	// 获取文章数据
	var article = new(Article)
	db := database.MysqlDb.First(article, articleId)
	if db.Error != nil {
		return db.Error
	}
	// 清除 label 数据
	err := ArticleLabelManager.ClearByArticleId(article.ArticleId)
	if err != nil {
		return err
	}
	// 添加标签
	for _, labelId := range labelIds {
		err := ArticleLabelManager.AddLabel(article.ArticleId, labelId)
		if err != nil {
			return err
		}
	}
	// 更新数据
	article.Title = title
	article.Content = content
	article.UpdateTime = time.Now()
	database.MysqlDb.Model(&article).Updates(article)
	return nil
}

// GetArticleInfoList 获取一列 articleInfo 数据
func (articleDao ArticleDao) GetArticleInfoList(labelIds []uint, preArticleId uint, limit uint) ([]ArticleInfo, error) {
	// 初始化数据
	limit = util.If(limit > 50, 50, limit).(uint)

	var articles []Article
	database.MysqlDb.Joins("join article_label al on article.article_id = al.article_id and  label_id in (?)", labelIds).Order("article_id DESC").Where("article.article_id > ?", preArticleId).Limit(int(limit)).Find(&articles)
	if len(articles) == 0 {
		return nil, errors.New("没有更多数据了")
	}
	var articleInfos []ArticleInfo
	for _, article := range articles {
		articleInfo, err := article.GetArticleInfo()
		if err != nil {
			return nil, err
		}
		articleInfos = append(articleInfos, *articleInfo)
	}
	return articleInfos, nil
}
