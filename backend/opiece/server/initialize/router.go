package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//_ "opiece/docs"
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
			//AllowAllOrigins:  true,
			AllowOrigins:     global.OConfig.Cors.Origins,
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			MaxAge:           12 * time.Hour,
		}))
	publicRouter := router.Group("v1")
	{
		router2.InitUserRouter(publicRouter)
		router2.InitPublicRouter(publicRouter)
		router2.InitImageRouter(publicRouter)
	}
	privateRouter := router.Group("v1")
	privateRouter.Use(middleware.JWTAuthMiddleware())
	{
		router2.InitArticleRouter(privateRouter)
		router2.InitUploadRouter(privateRouter)
		router2.InitDashboardRouter(privateRouter)
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
