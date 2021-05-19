package service

import (
	"errors"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
	"opiece/utils"
)

func Register(u model.User) (model.User, error) {
	var user model.User
	if !errors.Is(global.ODB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("用户名已存在")
	}
	u.Password = utils.MD5([]byte(u.Password))
	return u, global.ODB.Create(&u).Error
}

func Login(u *model.User) (*model.User, error) {
	var user model.User
	u.Password = utils.MD5([]byte(u.Password))
	err := global.ODB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return &user, err
}

func ChangePassword(user model.User, newPassword string) (error, model.User) {
	return nil, model.User{}
}
