package repository

import (
	"github.com/talhaunal7/expense-tracker/server/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) error
	FindUserByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db: db,
	}
}

func (userRepo userRepository) CreateUser(user entity.User) error {
	result := userRepo.db.Create(&user)
	return result.Error
}

func (userRepo userRepository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User
	result := userRepo.db.First("email = ?", email).First(&user)
	return user, result.Error
}
