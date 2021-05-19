package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"opiece/server/global"
	"opiece/server/model"
	"opiece/server/service"
	"opiece/server/service/response"
	"opiece/utils"
)

// PostArticle 新建文章接口
// @Summary 新建文章接口
// @Description 按照一定规则新建文章
// @Tags 文章管理相关接口
// @Accept application/json
// @Produce application/json
// @Param auth-token header string true "JWT用户令牌"
// @Param object body model.Article true "文章结构"
// @Router /article/post_article [post]
func PostArticle(c *gin.Context) {
	var A model.Article
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.FailWithDetailed(nil, "后端获取参数失败", c)
		c.Abort()
		return
	}
	err = json.Unmarshal(data, &A)
	if err != nil {
		response.FailWithDetailed(nil, "传入的参数格式有误", c)
		c.Abort()
		return
	}
	if A.Title == "" || A.Content == "" {
		response.FailWithDetailed(nil, "数据为空", c)
		c.Abort()
		return
	}
	if A.ContentHash == "" {
		A.ContentHash = utils.MD5([]byte(A.Content))
	}
	if A.BackgroundVisible {
		if A.BackgroundPic == "" && !A.BackgroundRandom {
			response.FailWithDetailed(&A, "缺少背景图片", c)
			c.Abort()
			return
		}
		// TODO: 从图片池中随机获取一个背景
	}
	articleRet, err := service.PostArticle(A)
	if err != nil {
		response.FailWithDetailed(articleRet, "插入数据失败", c)
	} else if articleRet != nil {
		response.FailWithDetailed(articleRet, "文章存在重复", c)
	} else {
		response.OkWithDetailed(articleRet, "文章插入成功", c)
	}
}

// UploadImage 图片上传接口
// @Summary 图片上传接口
// @Description 图片上传接口 单纯只接收图片
// @Tags 上传管理相关接口
// @Accept mpfd
// @Param auth-token header string true "JWT用户令牌"
// @Param file_hash formData string false "图片hash 为空则后端进行计算"
// @Param local_file formData file true "图片"
// @Router /upload/pic [post]
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
		response.FailWithDetailed(nil, "不支持的类型", c)
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
	err = service.UploadImage(Image)
	if err != nil {
		response.FailWithDetailed(nil, "上传图片失败: "+err.Error(), c)
		c.Abort()
		return
	}
	response.Ok(c)
}
