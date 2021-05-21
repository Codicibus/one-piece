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
		ArticleRouter.POST("post_article", v1.PostArticle)
		ArticleRouter.POST("remove_article", v1.RemoveArticle)
	}
}
