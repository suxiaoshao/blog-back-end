package router

import (
	"github.com/gin-gonic/gin"
	"nextBlogServer/controllers"
	"nextBlogServer/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	visitor := r.Group("/api/visitor")
	{
		visitor.GET("/articleNum", controllers.ArticleNum)
		visitor.GET("/articleList", controllers.ArticleList)
		visitor.GET("/wallpaper/:pid", controllers.Wallpaper)
		visitor.GET("/article/:aid", controllers.GetArticle)
		visitor.GET("/typeList", controllers.TypeList)
		visitor.GET("/replyNum/:aid", controllers.RePlyNum)
		visitor.GET("/replyList/:aid", controllers.ReplyList)
		visitor.POST("/uploadReply", controllers.UploadReply)
	}
	/* 登入 */
	r.POST("/api/admin/login", controllers.Login)
	admin := r.Group("/api/admin", middleware.CheckLogin())
	{
		admin.GET("/checkLogin", func(context *gin.Context) {
			context.JSON(200, gin.H{})
		})
		admin.POST("/article", controllers.UpdateArticle)
	}
	return r
}
