package model

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"math/rand"
	"net/http"
)

var Client *mongo.Client
var articleDatabase *mongo.Collection
var replyDatabase *mongo.Collection

type ImageItem struct {
	Binary []byte `bson:"binary" json:"binary"`
	Url    string `bson:"url" json:"url"`
}


type LoginInfo struct {
	Password string `bson:"password" json:"password" form:"password"`
	User     string `bson:"user" json:"user" form:"user"`
}

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, _ = mongo.Connect(context.TODO(), clientOptions)
	articleDatabase = Client.Database("my_blog").Collection("article")
	replyDatabase = Client.Database("my_blog").Collection("reply")
}

func GetRandomImage() (string, error) {
	imageDatabase := Client.Database("my_blog").Collection("img")
	count, err := imageDatabase.CountDocuments(context.TODO(), gin.H{"flag_num": gin.H{"$exists": true}})
	if err != nil {
		return "", err
	}
	num := rand.Int63n(count-1) + 1
	image := ImageItem{make([]byte, 0), ""}
	err = imageDatabase.FindOne(context.TODO(), gin.H{"flag_num": num}).Decode(&image)
	if err != nil {
		return "", err
	}
	if len(image.Binary) == 0 {
		resp, _ := http.Get(image.Url)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		_, err = imageDatabase.UpdateOne(context.TODO(), gin.H{"flag_num": num}, gin.H{"$set": gin.H{"binary": content}, "$unset": gin.H{"done": ""}})
		if err != nil {
			fmt.Println(err)
		}
		return string(content), nil
	}
	return string(image.Binary), nil
}

func GetTypeList() ([]string, error) {
	infoDatabase := Client.Database("my_blog").Collection("info")
	result := struct {
		Type []string `json:"type" bson:"type"`
	}{}
	err := infoDatabase.FindOne(context.TODO(), gin.H{}).Decode(&result)
	if err != nil {
		return make([]string, 0), err
	}
	return result.Type, nil
}
