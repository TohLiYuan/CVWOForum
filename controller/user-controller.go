package controller

import (
	"App/data"
	"App/dto"
	"App/entity"
	"App/service"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context) error
	Login(ctx *gin.Context) (string, error)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(service service.UserService, jWtService service.JWTService) UserController {
	return &userController{
		userService: service,
		jwtService:  jWtService,
	}
}

func (c *userController) Register(ctx *gin.Context) error {
	var newUser entity.Users
	var err error

	err = ctx.ShouldBindJSON(&newUser)
	if err != nil {
		return err
	}

	err = c.userService.Register(newUser)
	if err != nil {
		return err
	}

	return nil

}

func (c *userController) Login(ctx *gin.Context) (string, error) {
	var credentials dto.SecureCredentials
	var err error

	err = ctx.ShouldBindJSON(&credentials)
	if err != nil {
		return "", err
	}

	isAuthenticated := c.userService.Login(credentials.Email, credentials.Password)
	if isAuthenticated {
		var user entity.Users
		err := data.DB.Where("email = ?", credentials.Email).First(&user).Error
		if err != nil {
			return "", err
		}
		return c.jwtService.Generate(user.UserID), nil
	}
	return "", errors.New("Incorrect email or password")
}
