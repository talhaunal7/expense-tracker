package service

import (
	"github.com/talhaunal7/expense-tracker/server/model/dto"
	"github.com/talhaunal7/expense-tracker/server/model/request"
)

type UserService interface {
	Register(*request.UserRegister) error
	Login(*request.UserLogin) (*dto.UserDto, *string, error)
	Logout(id float64) error
}
