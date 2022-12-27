package service

import (
	"App/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(entity.Users) entity.Users
	Login(entity.Users) entity.Users
}

func Register(newUser entity.Users) error {
	passHash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return err
	}
	newSecureUser := entity.SecureUsers{
		Email:    newUser.Email,
		Password: passHash,
		Username: newUser.Username,
	}

	db.Create(&newSecureUser)

	return nil
}

func getPasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 0)
	return string(hash), err
}
