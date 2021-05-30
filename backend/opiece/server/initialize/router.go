package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "opiece/docs"
	"opiece/server/global"
	"opiece/server/middleware"
	router2 "opiece/server/router"
)

func Routers() *gin.Engine {
	gin.SetMode(global.OConfig.Gin.Mode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	publicRouter := router.Group("v1")
	{
		router2.InitUserRouter(publicRouter)
		router2.InitDashboardRouter(publicRouter)
	}
	privateRouter := router.Group("v1")
	privateRouter.Use(middleware.JWTAuthMiddleware())
	{
		router2.InitArticleRouter(privateRouter)
		router2.InitUploadRouter(privateRouter)
		router2.InitImageRouter(privateRouter)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
