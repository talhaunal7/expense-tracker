package service

import "github.com/talhaunal7/expense-tracker/server/model"

type UserService interface {
	Register(*model.UserRegister) error
	Login(*model.UserLogin) (*model.UserDto, *string, error)
	Logout(id float64) error
}
