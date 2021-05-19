package service

import (
	"errors"
	"opiece/server/global"
	"opiece/server/model"
)

func PostArticle(article model.Article) error {
	var art model.Article
	err := global.ODB.Where("content = ? AND removed = false", article.Content).First(&art).Error
	// 如果该博客内容已经存在就不再上传
	if err == nil {
		return errors.New("文章已存在")
	}
	return global.ODB.Create(&article).Error
}

func RemoveArticle(article model.Article, force bool) error {
	return nil
}
