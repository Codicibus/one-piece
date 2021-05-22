package v1

import (
	"github.com/gin-gonic/gin"
	"opiece/server/service"
	"opiece/server/service/response"
)

func ImagePool(c *gin.Context) {
	response.OkWithDetailed(service.GetRandomBackgroundPic(), "", c)
}

func GetImage(c *gin.Context) {
	imageHash := c.Request.FormValue("image_hash")
	pic := service.GetBackgroundPicByHash(imageHash)
	if pic != nil {
		_, _ = c.Writer.Write(pic.ImageBin)
	} else {
		response.Fail(c)
	}
}
