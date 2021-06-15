package model

import (
	"github.com/google/uuid"
	"opiece/server/global"
)

type User struct {
	global.OModel
	Username string    `gorm:"username" json:"username,omitempty" form:"username" binding:"required" gorm:"comment:用户名"`
	Password string    `gorm:"password" json:"password,omitempty" form:"password" binding:"required" gorm:"comment:用户密码"`
	Email    string    `gorm:"email" json:"-" form:"email" gorm:"comment:用户邮箱"`
	UUID     uuid.UUID `gorm:"uuid" json:"-" gorm:"comment:用户UUID"`
}
