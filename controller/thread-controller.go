package controller

import (
	"App/entity"
	"App/service"

	"github.com/gin-gonic/gin"
)

type ThreadController interface {
	List() []entity.Threads
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.ThreadService
}

func New(service service.ThreadService) ThreadController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var t entity.Threads
	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		return err
	}
	c.service.Save(t)
	return nil
}

func (c *controller) List() []entity.Threads {
	return c.service.List()
}
