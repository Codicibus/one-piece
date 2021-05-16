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
