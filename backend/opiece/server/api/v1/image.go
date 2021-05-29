package v1

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"opiece/server/global"
	"opiece/server/model"
	"opiece/server/service"
	"opiece/server/service/response"
	"opiece/server/utils"
)

// ImagePool 图片池
// @Tags 图片接口
// @Summary 获取随机图片
// @Success 200 {object} model.PICs
// @Router /v1/image/random [get]
func ImagePool(c *gin.Context) {
	response.OkWithDetailed(service.GetRandomBackgroundPic(), "", c)
}

// GetImage 获取图片
// @Tags 图片接口
// @Summary 根据图片hash获取图片
// @Router /v1/image/query [get]
func GetImage(c *gin.Context) {
	imageHash := c.Request.FormValue("image_hash")
	pic := service.GetBackgroundPicByHash(imageHash)
	if pic != nil {
		_, _ = c.Writer.Write(pic.ImageBin)
	} else {
		response.Fail(c)
	}
}


// UploadImage 图片上传接口
// @Summary 图片上传接口
// @Description 图片上传接口 单纯只接收图片
// @Tags 上传管理相关接口
// @Accept mpfd
// @Param Authorization header string true "JWT用户令牌"
// @Param file_hash formData string false "图片hash 为空则后端进行计算"
// @Param local_file formData file true "图片"
// @success 200 {object} response.MsgResponse{data=string} "成功"
// @Router /v1/upload/pic [post]
func UploadImage(c *gin.Context) {
	var Image model.PICs
	fileHash := c.Request.FormValue("file_hash")
	file, header, err := c.Request.FormFile("local_file")
	if err != nil {
		response.FailWithDetailed(nil, "文件有误: "+err.Error(), c)
		c.Abort()
		return
	}
	mime := header.Header.Get("Content-Type")
	if !utils.Target(mime).In(global.MIMEImages) {
		response.FailWithDetailed(nil, "不支持的文件类型", c)
		c.Abort()
		return
	}
	Image.ImageMIMEType = mime
	Image.ImageSize = header.Size
	Image.ImageName = header.Filename
	Image.ImageBin, err = ioutil.ReadAll(file)
	if err != nil {
		response.FailWithDetailed(nil, "读取文件错误: "+err.Error(), c)
		c.Abort()
		return
	}
	Image.ImageHash = fileHash
	if Image.ImageHash == "" {
		Image.ImageHash = utils.MD5(Image.ImageBin)
	}
	Image.ImagePath = "binary"
	imageURI, err := service.UploadImage(Image)
	if err != nil && imageURI == "" {
		response.FailWithDetailed(nil, "上传图片失败: "+err.Error(), c)
		c.Abort()
		return
	} else if err != nil {
		response.OkWithDetailed(imageURI, err.Error(), c)
	} else {
		response.OkWithDetailed(imageURI, "图片上传成功！", c)
	}
}