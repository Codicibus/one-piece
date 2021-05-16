package model

import "opiece/server/global"

type User struct {
	global.OModel
	Username string `gorm:"username" json:"username" form:"username" binding:"required" gorm:"comment:用户名"`
	Password string `gorm:"password" json:"password" form:"password" binding:"required" gorm:"comment:用户密码"`
	Email string `gorm:"email" json:"email" form:"email" gorm:"comment:用户邮箱"`
}
