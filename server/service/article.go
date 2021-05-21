package service

import (
	"errors"
	"fmt"
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
	var Article model.Article
	fmt.Println(article)
	err := global.ODB.Where("removed = false AND content_hash = ?", article.ContentHash).First(&Article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在！")
		}
		return err
	}
	// 执行记录删除指令
	if !force {
		return global.ODB.Where("content_hash = ?", article.ContentHash).Delete(&article).Error
	}
	return global.ODB.Model(&article).Where("content_hash = ?", article.ContentHash).Update("removed", true).Error
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
