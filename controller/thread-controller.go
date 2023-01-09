package controller

import (
	"App/data"
	"App/entity"
	"App/middleware"
	"App/service"
	"App/validators"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ThreadController interface {
	List(ctx *gin.Context) []entity.Threads
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	GetPost(ctx *gin.Context) (entity.Threads, error)
}

type threadController struct {
	threadService service.ThreadService
}

var validate *validator.Validate

func NewThreadController(service service.ThreadService) ThreadController {
	validate = validator.New()
	validate.RegisterValidation("check-category", validators.ValidateCategory)

	return &threadController{
		threadService: service,
	}
}

func (c *threadController) Save(ctx *gin.Context) error {
	var t entity.Threads
	var u entity.Users
	data.DB.Where("user_id = ?", middleware.Uid).First(&u)
	var lastEntry entity.Threads
	data.DB.Table("threads").Last(&lastEntry)
	t.ID = lastEntry.ID + 1
	t.URL = fmt.Sprint("https://www.cvwoforums.com/threads/", lastEntry.ID+1)

	t.Users = u
	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		return err
	}
	err = validate.Struct(t)
	if err != nil {
		return err
	}
	c.threadService.Save(t)
	return nil
}

func (c *threadController) Update(ctx *gin.Context) error {
	var t entity.Threads
	var u entity.Users

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	t.ID = id

	// Validate user performing update action
	if t.Uid != middleware.Uid {
		return errors.New("Unauthorised access")
	}

	data.DB.Where("user_id = ?", middleware.Uid).First(&u)
	t.Users = u

	err = ctx.ShouldBindJSON(&t)
	if err != nil {
		return err
	}

	c.threadService.Update(t)
	return nil
}

func (c *threadController) Delete(ctx *gin.Context) error {
	var t entity.Threads
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	data.DB.Where("ID = ?", id).First(&t)

	// Validate user performing delete action
	if t.Uid != middleware.Uid {
		return errors.New("Unauthorised access")
	}
	c.threadService.Delete(t)
	return nil
}

func (c *threadController) List(ctx *gin.Context) []entity.Threads {
	category := ctx.Param("category")

	return c.threadService.List(category)
}

func (c *threadController) GetPost(ctx *gin.Context) (entity.Threads, error) {
	var t entity.Threads
	id, err := strconv.ParseUint(ctx.Param("tid"), 0, 0)
	if err != nil {
		return t, err
	}
	t = c.threadService.GetPost(id)
	return t, nil
}
