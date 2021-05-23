package service

import (
	"errors"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
)

func UploadImage(image model.PICs) (string, error) {
	var Image model.PICs
	err := global.ODB.Where("image_hash = ?", image.ImageHash).First(&Image).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = global.ODB.Create(&image).Error
			if err != nil {
				return "", err
			}
			return "/api/image/query?image_hash=" + image.ImageHash, nil
		}
	}
	return "/api/image/query?image_hash=" + Image.ImageHash, errors.New("图片已经存在")
}
