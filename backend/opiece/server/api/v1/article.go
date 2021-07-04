package v1

import (
	"encoding/json"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"opiece/server/middleware"
	"opiece/server/model"
	"opiece/server/service"
	"opiece/server/service/response"
	"opiece/server/utils"
	"strconv"
	"strings"
)

// PostArticle 新增文章接口
// @Summary 新增文章接口
// @Description 按照一定规则新建文章
// @Tags 文章管理相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "JWT用户令牌"
// @Param title formData string true "文章标题"
// @Param content formData string true "文章内容"
// @Param content_hash formData string false "文章hash值，为空则后端计算"
// @Param background_visible formData bool false "是否显示背景图，默认false，为true时必须选择设置random和pic其中之一"
// @Param background_random formData bool false "当background_visible为true时"
// @Param background_pic formData string false "当background_visible为true时"
// @Router /v1/article/post [post]
// @Response 200 {object} model.Article
func PostArticle(c *gin.Context) {
	var A model.Article
	A.Title = c.Request.FormValue("title")
	A.Content = c.Request.FormValue("content")
	A.BackgroundVisible = utils.GetBool(c.Request.FormValue("background_visible"))
	A.BackgroundRandom = utils.GetBool(c.Request.FormValue("background_random"))
	A.BackgroundPic = c.Request.FormValue("background_pic")
	A.Category = c.Request.FormValue("category")
	A.ContentHash = c.Request.FormValue("content_hash")
	//A.UUID = c.Request.Header.Get("Authorization")
	jwt := middleware.NewJWT()
	claims, err := jwt.ParseToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		response.FailWithDetailed(claims, "解析token失败："+err.Error(), c)
		c.Abort()
		return
	}
	A.UUID = claims.UUID
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
			A.BackgroundPic = "/v1/image/query?image_hash=" + service.GetRandomBackgroundPic().ImageHash
		}
	}
	if !strings.HasPrefix(A.BackgroundPic, "http") {
		A.BackgroundPic = "v1/image/query?image_hash=" + A.BackgroundPic
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

func UpdateArticle(c *gin.Context) {
	oldArticleHash := c.Request.FormValue("old_article_hash")
	jwt := middleware.NewJWT()
	claims, err := jwt.ParseToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		response.FailWithDetailed(claims, "解析token失败："+err.Error(), c)
		c.Abort()
		return
	}
	oldArticleUUID := claims.UUID
	oldArticle := model.Article{
		ContentHash: oldArticleHash,
		UUID:        oldArticleUUID,
	}
	var A model.Article
	A.UUID = claims.UUID
	A.Title = c.Request.FormValue("title")
	A.Content = c.Request.FormValue("content")
	A.BackgroundVisible = utils.GetBool(c.Request.FormValue("background_visible"))
	A.BackgroundRandom = utils.GetBool(c.Request.FormValue("background_random"))
	A.BackgroundPic = c.Request.FormValue("background_pic")
	A.Category = c.Request.FormValue("category")
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
			A.BackgroundPic = "/v1/image/query?image_hash=" + service.GetRandomBackgroundPic().ImageHash
		}
	}
	if !strings.HasPrefix(A.BackgroundPic, "http") {
		A.BackgroundPic = "v1/image/query?image_hash=" + A.BackgroundPic
	}
	article, err := service.UpdateArticle(oldArticle, A)
	if err != nil {
		response.FailWithDetailed(article, err.Error(), c)
		c.Abort()
		return
	}
	response.OkWithDetailed(article, "操作成功", c)
}

// RemoveArticle 文章删除接口
// @Summary 文章删除接口
// @Description 文章删除接口 接收一个文章对象
// @Tags 文章管理相关接口
// @Accept json
// @Param Authorization header string true "JWT用户令牌"
// @Param article body model.Article true "文章对象"
// @Router /v1/article/remove [post]
func RemoveArticle(c *gin.Context) {
	body := c.Request.Body
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
	Article.UUID, _ = utils.GetUUIDFromToken(c.Request.Header.Get("Authorization"))
	err = service.RemoveArticle(Article)
	if err != nil {
		response.FailWithDetailed(nil, "删除失败: "+err.Error(), c)
		return
	}
	response.Ok(c)
}

// DeleteArticle delete article by content hash
func DeleteArticle(c *gin.Context) {
	contentHash := c.Request.FormValue("content_hash")
	if contentHash == "" {
		response.FailWithDetailed(nil, "未指定文章", c)
		c.Abort()
		return
	}
	uuid, _ := utils.GetUUIDFromToken(c.Request.Header.Get("Authorization"))
	if err := service.DeleteArticle(contentHash, uuid.String()); err != nil {
		response.FailWithDetailed(nil, "操作失败: "+err.Error(), c)
		c.Abort()
		return
	}
	response.Ok(c)
}

// GetArticle 文章获取接口
// @Summary 文章获取接口
// @Description 文章获取接口
// @Tags 文章管理相关接口
// @Accept json
// @Param page_size query string false "请求页数"
// @Param page_num query string false "请求页码"
// @Router /v1/article/get [get]
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	articles, err := service.GetArticles(pageSize, pageNum)
	if err != nil {
		response.FailWithDetailed(nil, "获取文章失败："+err.Error(), c)
		c.Abort()
		return
	}
	count, err := service.GetAllArticlesCount()
	if err != nil {
		count = -1
		// TODO: record log
	}
	response.OkWithDetailed(map[string]interface{}{
		"page_size": pageSize,
		"page_num":  pageNum,
		"total":     count,
		"articles":  articles,
	}, "操作成功！", c)
}

// GetArticleByCategory 按分类获取文章
// @Tags 文章管理相关接口
// @Param Authorization header string true "JWT用户令牌"
// @Param name query string false "分类名称"
// @Router /v1/article/category [get]
func GetArticleByCategory(c *gin.Context) {
	category := c.Query("name")
	if category == "" {
		response.FailWithDetailed(nil, "缺少category参数", c)
		c.Abort()
		return
	}
	articles, err := service.GetArticlesByCategory(category)
	if err != nil {
		response.FailWithDetailed(articles, "获取文章失败: "+err.Error(), c)
		c.Abort()
		return
	}
	response.OkWithDetailed(articles, "操作成功", c)
}

// ImportArticleFromFile
// 从前端表单接收多个文件，并读取内容加以处理存入数据库
func ImportArticleFromFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithDetailed(form, "操作失败: "+err.Error(), c)
		c.Abort()
		return
	}
	genMsg := func(filename, msg string) map[string]string {
		return map[string]string{
			"file_name": filename,
			"msg":       msg,
		}
	}
	var bf model.BinaryFile
	failed := make([]map[string]string, 0)
	for _, file := range form.File {
		for _, f := range file {
			fd, err := f.Open()
			if err != nil {
				failed = append(failed, genMsg(f.Filename, err.Error()))
				continue
			}
			data, _ := ioutil.ReadAll(fd)
			bf.FileName = f.Filename
			bf.FileSize = int(f.Size)
			bf.FileData = data
			bf.FileHash = utils.MD5(data)
			mtype := mimetype.Detect(data)
			if !strings.Contains(mtype.String(), "text/plain") {
				failed = append(failed, genMsg(f.Filename, "非法文件类型"))
				continue
			}
			err = service.UploadBinaryFile(bf)
			if err != nil {
				failed = append(failed, genMsg(f.Filename, err.Error()))
				continue
			}
		}
	}
	if len(failed) != 0 {
		response.FailWithDetailed(failed, "操作有错误", c)
		return
	}
	response.Ok(c)
}
