package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	articleDatabase := client.Database("my_blog").Collection("article")
	replyDatabase := client.Database("my_blog").Collection("reply")
	//取消目录和timeStr
	_, _ = articleDatabase.UpdateMany(context.TODO(), gin.H{}, gin.H{"$unset": gin.H{"directory": "", "time_str": ""}})
	_, _ = replyDatabase.UpdateMany(context.TODO(), gin.H{}, gin.H{"$unset": gin.H{"time_str": ""}})
	infoDatabase := client.Database("my_blog").Collection("info")

	//更新类型数组
	_, _ = infoDatabase.UpdateOne(context.TODO(), gin.H{}, gin.H{"$set": gin.H{"type": []string{"学习", "代码", "其他", "工具"}}})

	//文章的更新
	cursor, _ := articleDatabase.Find(context.TODO(), gin.H{})
	for cursor.Next(context.TODO()) {
		result := struct {
			Aid  int64 `bson:"aid" json:"aid"`
			Type int64 `bson:"type" json:"type"`
		}{Type: -1, Aid: -1}
		err := cursor.Decode(&result)

		if result.Aid == -1 {
			fmt.Println(err, result)
			continue
		}
		//更换文章数组
		if result.Type != -1 {
			_, _ = articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": result.Aid}, gin.H{"$set": gin.H{"type": []int64{result.Type}}})
		}
	}
	cursor, _ = articleDatabase.Find(context.TODO(), gin.H{})
	for cursor.Next(context.TODO()) {
		result := struct {
			Aid       int64   `bson:"aid" json:"aid"`
			TimeStamp float64 `bson:"time_stamp" json:"timeStamp"`
		}{Aid: -1, TimeStamp: -1}
		err := cursor.Decode(&result)

		if result.Aid == -1 {
			fmt.Println(err, result)
			continue
		}

		/*if result.TimeStamp != -1 {
			if result.TimeStamp < 158978171 {
				_, _ = articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": result.Aid}, gin.H{"$set": gin.H{"time_stamp": int64(result.TimeStamp * 1000000)}})
			} else {
				_, _ = articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": result.Aid}, gin.H{"$set": gin.H{"time_stamp": int64(result.TimeStamp * 1000)}})
			}
		}*/

		//更新评论总数
		count, err := replyDatabase.CountDocuments(context.TODO(), gin.H{"aid": result.Aid})
		if err != nil {
			continue
		}
		_, _ = articleDatabase.UpdateOne(context.TODO(), gin.H{"aid": result.Aid}, gin.H{"$set": gin.H{"reply_num": count}})
	}

	//	评论的更新
	cursor, _ = replyDatabase.Find(context.TODO(), gin.H{})
	for cursor.Next(context.TODO()) {
		result := struct {
			Rid       int64   `bson:"rid" json:"rid"`
			TimeStamp float64 `bson:"time_stamp" json:"timeStamp"`
		}{Rid: -1, TimeStamp: -1}
		err := cursor.Decode(&result)

		if result.Rid == -1 {
			fmt.Println(err, result)
			continue
		}
		/*if result.TimeStamp != -1 {
			_, _ = replyDatabase.UpdateOne(context.TODO(), gin.H{"rid": result.Rid}, gin.H{"$set": gin.H{"time_stamp": int64(result.TimeStamp * 1000)}})
		}*/
	}
}
