package router

import (
	"github.com/gin-gonic/gin"
	v1 "opiece/server/api/v1"
)

func InitImageRouter(Router *gin.RouterGroup) {
	ImageRouter := Router.Group("image")
	{
		ImageRouter.GET("/random", v1.ImagePool)
		ImageRouter.GET("/query", v1.GetImage)
	}
}
