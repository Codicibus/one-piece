package router

import (
	"github.com/gin-gonic/gin"
	v1 "opiece/server/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", v1.Register)
		UserRouter.POST("login", v1.Login)
	}
}

func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("article")
	{
		ArticleRouter.POST("post", v1.PostArticle)
		ArticleRouter.POST("remove", v1.RemoveArticle)
		ArticleRouter.GET("get", v1.GetArticle)
	}
}

func InitUploadRouter(Router *gin.RouterGroup) {
	UploadRouter := Router.Group("upload")
	{
		UploadRouter.POST("/pic", v1.UploadImage)
	}
}

func InitImageRouter(Router *gin.RouterGroup) {
	ImageRouter := Router.Group("image")
	{
		ImageRouter.GET("/random", v1.ImagePool)
		ImageRouter.GET("/query", v1.GetImage)
	}
}