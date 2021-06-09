package service

import (
	"errors"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
)

func UploadBinaryFile(BinaryFile model.BinaryFile) error {
	var File model.BinaryFile
	err := global.ODB.Where("file_hash = ?", BinaryFile.FileHash).First(&File).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// TODO insert the record into database
			err = global.ODB.Create(&BinaryFile).Error
			if err != nil {
				return  err
			}
			return nil
		}
	}
	return errors.New("文件已经存在")
}

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
