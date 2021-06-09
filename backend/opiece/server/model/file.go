package model

import "opiece/server/global"

type BinaryFile struct {
	global.OModel
	FileName string `gorm:"file_name"`
	FileData []byte `gorm:"file_data"`
	FileHash string `gorm:"file_hash"`
	FileSize int    `gorm:"file_size"`
}
