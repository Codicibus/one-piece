package initialize

import (
	"github.com/gin-gonic/gin"
	"opiece/server/global"
	router2 "opiece/server/router"
)

func Routers() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	privateRouter := router.Group("")
	router2.InitUserRouter(privateRouter)
	privateRouter.Use(global.OAuthJWT.MiddlewareFunc())
	return router
}