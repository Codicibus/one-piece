package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "opiece/docs"
	"opiece/server/global"
	"opiece/server/middleware"
	router2 "opiece/server/router"
	"time"
)

func Routers() *gin.Engine {
	gin.SetMode(global.OConfig.Gin.Mode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowCredentials: true,
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			MaxAge: 12 * time.Hour,
		}))
	publicRouter := router.Group("v1")
	{
		router2.InitUserRouter(publicRouter)
	}
	privateRouter := router.Group("v1")
	privateRouter.Use(middleware.JWTAuthMiddleware())
	{
		router2.InitArticleRouter(privateRouter)
		router2.InitUploadRouter(privateRouter)
		router2.InitImageRouter(privateRouter)
		router2.InitDashboardRouter(privateRouter)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
