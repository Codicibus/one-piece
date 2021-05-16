package service

import "opiece/server/model"

func Register(user model.User) (error, model.User) {
	return nil, model.User{}
}

func Login(user model.User) error {
	return nil
}

func ChangePassword(user model.User, newPassword string) (error, model.User) {
	return nil, model.User{}
}
