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

// PostArticle 新增文章接口
// @Summary 新增文章接口
// @Description 按照一定规则新建文章
// @Tags 文章管理相关接口
// @Accept application/json
// @Produce application/json
// @Param auth-token header string true "JWT用户令牌"
// @Param title formData string true "文章标题"
// @Param content formData string true "文章内容"
// @Param content_hash formData string false "文章hash值，为空则后端计算"
// @Param background_visible formData bool false "是否显示背景图，默认false，为true时必须选择设置random和pic其中之一"
// @Param background_random formData bool false "当background_visible为true时"
// @Param background_pic formData string false "当background_visible为true时"
// @Router /article/post_article [post]
func PostArticle(c *gin.Context) {
	var A model.Article
	A.Title = c.Request.FormValue("title")
	A.Content = c.Request.FormValue("content")
	A.BackgroundVisible = utils.GetBool(c.Request.FormValue("background_visible"))
	A.BackgroundRandom = utils.GetBool(c.Request.FormValue("background_random"))
	A.BackgroundPic = c.Request.FormValue("background_pic")
	A.ContentHash = c.Request.FormValue("content_hash")
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
		if A.BackgroundPic == "" && A.BackgroundRandom {
			A.BackgroundPic = "/api/image/query?image_hash=" + service.GetRandomBackgroundPic().ImageHash
		}
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

// RemoveArticle 文章删除接口
// @Summary 文章删除接口
// @Description 文章删除接口 接收一个文章对象
// @Tags 文章管理相关接口
// @Accept json
// @Param auth-token header string true "JWT用户令牌"
// @Param file_hash body model.Article true "文章对象"
// @Router /article/remove_article [post]
func RemoveArticle(c *gin.Context) {
	body := c.Request.Body
	force := utils.GetBool(c.Request.PostFormValue("force"))
	var Article model.Article
	data, err := ioutil.ReadAll(body)
	if err != nil {
		response.FailWithDetailed(nil, err.Error(), c)
		c.Abort()
		return
	}
	err = json.Unmarshal(data, &Article)
	if err != nil {
		response.FailWithDetailed(nil, err.Error(), c)
		c.Abort()
		return
	}
	err = service.RemoveArticle(Article, force)
	if err != nil {
		response.FailWithDetailed(nil, "删除失败: "+err.Error(), c)
		return
	}
	response.Ok(c)
}
