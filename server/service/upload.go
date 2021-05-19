package service

import (
	"errors"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
)

func UploadImage(image model.PICs) error {
	var Image model.PICs
	err := global.ODB.Where("image_hash = ?", image.ImageHash).First(&Image).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return global.ODB.Create(&image).Error
		}
	}
	return errors.New("图片已经存在")
}
