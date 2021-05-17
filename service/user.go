package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"opiece/server/global"
	"opiece/server/model"
)

func Register(u model.User) (model.User, error) {
	var user model.User
	if !errors.Is(global.ODB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("用户名已存在")
	}
	h := md5.New()
	h.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(h.Sum(nil))
	return u, global.ODB.Create(&u).Error
}

func Login(user model.User) error {
	return nil
}

func ChangePassword(user model.User, newPassword string) (error, model.User) {
	return nil, model.User{}
}
