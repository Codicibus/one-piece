package model

import "opiece/server/global"

type Article struct {
	global.OModel
	Removed     bool   `gorm:"removed" gorm:"comment:是否已经删除"`
	Title       string `gorm:"title" gorm:"comment:文章标题"`
	Content     string `gorm:"content" gorm:"comment:文章内容"`
	Category    string `gorm:"category" gorm:"comment:文章分类"`
	ContentHash string `gorm:"content_hash" gorm:"comment:文章内容的唯一值"`
}
