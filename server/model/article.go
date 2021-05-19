package model

import "opiece/server/global"

// Article 文章结构模型
type Article struct {
	global.OModel
	Removed           bool   `gorm:"removed" json:"-" gorm:"comment:是否已经删除" example:"false"`
	Title             string `gorm:"title" json:"title" gorm:"comment:文章标题" example:"世界，你好！"`
	Content           string `gorm:"content" json:"content" gorm:"comment:文章内容" example:"你好，世界！"`
	Category          string `gorm:"category" gorm:"default:default" json:"category" gorm:"comment:文章分类" example:"default"`
	ContentHash       string `gorm:"content_hash" json:"-" gorm:"comment:文章内容的唯一值"`
	BackgroundPic     string `gorm:"background_pic" json:"background_pic" gorm:"comment:文章背景图片" example:"http://example.com/a.png"`
	BackgroundRandom  bool   `gorm:"background_random" gorm:"default:false" json:"background_random" gorm:"comment:是否设置随机背景" example:"true"`
	BackgroundVisible bool   `gorm:"background_visible" gorm:"default:false" json:"background_visible" gorm:"comment:是否显示背景图" example:"true"`
}
