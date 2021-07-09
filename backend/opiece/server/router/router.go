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

func InitPublicRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("article")
	{
		ArticleRouter.GET("get", v1.GetArticle)
		ArticleRouter.GET("category", v1.GetArticleByCategory)
	}
}

func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("article")
	{
		ArticleRouter.POST("post", v1.PostArticle)
		ArticleRouter.POST("remove", v1.RemoveArticle)
		ArticleRouter.POST("delete", v1.DeleteArticle)
		ArticleRouter.POST("update", v1.UpdateArticle)
		ArticleRouter.POST("import_files", v1.ImportArticleFromFile)
		ArticleRouter.GET("export_file", v1.ExportArticleToZipFile)
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

func InitDashboardRouter(Router *gin.RouterGroup) {
	DashboardRouter := Router.Group("dashboard")
	{
		DashboardRouter.GET("/article_count", v1.GetArticleStatHttp)
	}
	Ws := DashboardRouter.Group("ws")
	{
		Ws.GET("/systat", v1.GetSysStat)
		Ws.GET("/article_count_ws", v1.GetArticleStat)
	}
}
