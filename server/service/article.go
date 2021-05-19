package service

import (
	"errors"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
)

func PostArticle(article model.Article) (*model.Article, error) {
	var art model.Article
	err := global.ODB.Preload("article").Where("content_hash = ?", article.ContentHash).First(&art).Error
	// 如果记录没被找到则插入
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, global.ODB.Create(&article).Error
	}
	return &art, nil
}

func RemoveArticle(article model.Article, force bool) error {
	return nil
}
