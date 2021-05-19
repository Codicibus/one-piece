package model

import "opiece/server/global"

// PICs 图片上传结构模型
type PICs struct {
	global.OModel
	ImageSize     int64  `gorm:"image_size" json:"image_size" gorm:"comment:图片大小" example:"1024"`
	ImageName     string `gorm:"image_name" json:"image_name" gorm:"comment:图片名称" example:"example.png"`
	ImagePath     string `gorm:"image_path" json:"image_path" gorm:"comment:图片路径/url路径" example:"http://exmaple.com/example.png"`
	ImageHash     string `gorm:"image_hash" json:"image_hash" gorm:"comment:图片唯一值" example:"6a7bd4895ed239a4dc257d962da8a3ce"`
	ImageMIMEType string `gorm:"image_mime_type" json:"image_mime_type" gorm:"comment:图片类型" example:"png"`
	ImageBin      []byte `gorm:"image_bin" json:"image_bin" gorm:"comment:图片二进制内容"`
}
