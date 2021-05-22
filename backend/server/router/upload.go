package router

import (
	"github.com/gin-gonic/gin"
	v1 "opiece/server/api/v1"
)

func InitUploadRouter(Router *gin.RouterGroup) {
	UploadRouter := Router.Group("upload")
	{
		UploadRouter.POST("/pic", v1.UploadImage)
	}
}
