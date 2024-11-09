package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/talhaunal7/expense-tracker/server/entity"
	"github.com/talhaunal7/expense-tracker/server/model/dto"
	"github.com/talhaunal7/expense-tracker/server/model/request"
	"github.com/talhaunal7/expense-tracker/server/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u *UserServiceImpl) Register(userRegisterRequest *request.UserRegister) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), 10)
	if err != nil {
		return err
	}
	user := entity.User{
		Email:     userRegisterRequest.Email,
		Password:  string(hash),
		FirstName: userRegisterRequest.FirstName,
		LastName:  userRegisterRequest.LastName,
	}
	if err = u.userRepository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) Login(userLoginRequest *request.UserLogin) (*dto.UserDto, *string, error) {

	user, err := u.userRepository.FindUserByEmail(userLoginRequest.Email)
	if err != nil {
		return nil, nil, err
	}

	if user.ID == 0 {
		return nil, nil, errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return nil, nil, errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, nil, err
	}

	userDto := dto.UserDto{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return &userDto, &tokenString, nil

}

func (u *UserServiceImpl) Logout(id float64) error {
	fmt.Println("logout test", id)
	/*userId := strconv.FormatFloat(id, 'f', 0, 64)
	err := u.redis.Remove(userId)
	if err != nil {
		return err
	}*/
	return nil
}
