package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "opiece/docs"
	"opiece/server/middleware"
	router2 "opiece/server/router"
)

func Routers() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	publicRouter := router.Group("api")
	{
		router2.InitUserRouter(publicRouter)
	}
	privateRouter := router.Group("api")
	privateRouter.Use(middleware.JWTAuthMiddleware())
	{
		router2.InitArticleRouter(privateRouter)
		router2.InitUploadRouter(privateRouter)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
