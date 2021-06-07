package router

import (
	"blogServer/controllers"
	"blogServer/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	visitor := r.Group("/api/visitor")
	{
		visitor.GET("/articleList", controllers.ArticleList)
		visitor.GET("/article/:aid", controllers.GetArticleDetail)
		visitor.GET("/replyList/:aid", controllers.ReplyList)
		visitor.POST("/uploadReply", controllers.UploadReply)
	}
	/* 登入 */
	r.POST("/api/admin/login", controllers.Login)
	r.POST("/api/admin/image",controllers.ImageUpload)
	admin := r.Group("/api/admin", middleware.CheckLogin())
	{
		admin.GET("/checkLogin", func(context *gin.Context) {
			context.JSON(200, gin.H{})
		})
		admin.POST("/article", controllers.UpdateArticle)
	}
	return r
}
