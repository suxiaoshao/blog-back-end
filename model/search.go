package model

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nextBlogServer/util"
)

type SearchQuery struct {
	Offset     int64
	Limit      int64
	SortNumber int64
	Query      interface{}
}

//
func GetSearchQuery(offset int64, limit int64, searchName string, typeList []int64, sortNumber int64, allType bool) *SearchQuery {
	typeQuery := util.If(allType, "$all", "$in").(string)
	return &SearchQuery{
		Offset:     util.If(offset < 0, 0, offset).(int64),
		Limit:      util.If(limit > 50 || limit < 0, 50, limit).(int64),
		SortNumber: util.If(sortNumber != 1 && sortNumber != -1, -1, sortNumber).(int64),
		Query: util.If(len(typeList) > 0,
			gin.H{"title": gin.H{"$regex": primitive.Regex{Pattern: searchName, Options: "i"}}, "type": gin.H{typeQuery: typeList}},
			gin.H{"title": gin.H{"$regex": primitive.Regex{Pattern: searchName, Options: "i"}}}),
	}
}

// 返回文章列表
func (searchQuery *SearchQuery) GetArticleList() ([]ArticleItem, error) {

	//舒适化对象
	Result := ArticleItem{
		Title:    "",
		Aid:      0,
		Type:     []int64{},
		ReplyNum: 0,
		ReadNum:  0,
	}

	//初始化切片
	resultList := make([]ArticleItem, 0)

	//生成 findOption
	findOption := options.Find()
	findOption.SetSort(gin.H{"aid": searchQuery.SortNumber})
	findOption.SetSkip(searchQuery.Offset)
	findOption.SetLimit(searchQuery.Limit)

	cur, err := articleDatabase.Find(context.TODO(), searchQuery.Query, findOption)
	if err != nil {
		return resultList, err
	}
	for cur.Next(context.TODO()) {
		result := Result
		err = cur.Decode(&result)
		if err == nil {
			resultList = append(resultList, result)
		}
	}
	return resultList, nil
}

//返回文章数
func (searchQuery *SearchQuery) GetArticleNum() (int64, error) {

	count, err := articleDatabase.CountDocuments(context.TODO(), searchQuery.Query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
