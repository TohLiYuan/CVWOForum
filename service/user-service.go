package service

import (
	"App/data"
	"App/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(newUser entity.Users) error
	Login(email string, password string) bool
}

type userService struct {
	userData data.Data
}

func NewUserService(data data.Data) UserService {
	return &userService{
		userData: data,
	}
}

func (service *userService) Register(newUser entity.Users) error {
	var storedUser entity.Users
	err := data.DB.Where("email = ?", newUser.Email).First(&storedUser).Error
	if err == nil {
		return errors.New("Email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	service.userData.SaveUser(newUser)
	return nil
}

func (service *userService) Login(email string, password string) bool {
	var storedUser entity.Users
	var err error
	err = data.DB.Where("email = ?", email).First(&storedUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
