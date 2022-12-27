package service

import (
	"App/entity"

	"gorm.io/gorm"
)

var db *gorm.DB

type ThreadService interface {
	Save(entity.Threads) entity.Threads
	List() []entity.Threads
}

type threadService struct {
	threads []entity.Threads
}

func NewThreadService() ThreadService {
	return &threadService{
		threads: []entity.Threads{},
	}
}

func (service *threadService) Save(thread entity.Threads) entity.Threads {
	db.Create(&thread)
	return thread
}

func (service *threadService) List() []entity.Threads {
	return service.threads
}
