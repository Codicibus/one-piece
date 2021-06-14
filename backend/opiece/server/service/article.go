package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
)

func UpdateArticle(oldArticle, newArticle model.Article) (*model.Article, error) {
	var art model.Article
	err := global.ODB.Preload("article").Where("content_hash = ? AND uuid = ?", oldArticle.ContentHash, oldArticle.UUID).First(&art).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("not found then create")
		if errors.Is(global.ODB.Preload("article").Where("content_hash = ? AND uuid = ?", newArticle.ContentHash, newArticle.UUID).First(&art).Error, gorm.ErrRecordNotFound) {
			return &newArticle, global.ODB.Create(&newArticle).Error
		}
		return &art, errors.New("该条记录已存在")
	} else {
		return &newArticle, global.ODB.Model(oldArticle).Omit("created_at", "deleted_at", "uuid", "id").Where("content_hash = ? AND uuid = ?", oldArticle.ContentHash, oldArticle.UUID).Select("*").Updates(newArticle).Error
	}
}

func PostArticle(article model.Article) (*model.Article, error) {
	var art model.Article
	err := global.ODB.Preload("article").Where("content_hash = ? AND uuid = ?", article.ContentHash, article.UUID).First(&art).Error
	// 如果记录没被找到则插入
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, global.ODB.Create(&article).Error
	}
	return &art, nil
}

func RemoveArticle(article model.Article) error {
	var Article model.Article
	err := global.ODB.Where("content_hash = ? AND uuid = ?", article.ContentHash, article.UUID).First(&Article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在！")
		}
		return err
	}
	return global.ODB.Where("content_hash = ? AND uuid = ?", article.ContentHash, article.UUID).Delete(&article).Error
}

func GetArticles(pageSize int, pageNum int) ([]*model.Article, error) {
	var articles []*model.Article
	return articles, global.ODB.Order("created_at desc").Limit(pageSize).Offset(pageNum).Find(&articles).Error
}

func GetArticlesByCategory(category string) ([]*model.Article, error) {
	var articles []*model.Article
	return articles, global.ODB.Where("category = ?", category).Find(&articles).Error
}

func GetBackgroundPicByHash(hash string) *model.PICs {
	var pic model.PICs
	err := global.ODB.Where("image_hash = ?", hash).First(&pic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// TODO: write the error log
			return nil
		}
		// TODO: write the error log
	}
	return &pic
}

func GetRandomBackgroundPic() *model.PICs {
	var pics model.PICs
	err := global.ODB.Order("random()").Limit(1).First(&pics).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &pics
	}
	return nil
}
