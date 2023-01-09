package service

import (
	"App/data"
	"App/entity"
)

type ThreadService interface {
	Save(entity.Threads) entity.Threads
	Update(thread entity.Threads) entity.Threads
	Delete(thread entity.Threads)
	List(category string) []entity.Threads
	GetPost(tid uint64) entity.Threads
}

type threadService struct {
	threadData data.Data
}

func NewThreadService(data data.Data) ThreadService {
	return &threadService{
		threadData: data,
	}
}

func (service *threadService) Save(thread entity.Threads) entity.Threads {
	service.threadData.Save(thread)
	return thread
}

func (service *threadService) Update(thread entity.Threads) entity.Threads {
	service.threadData.Update(thread)
	return thread
}

func (service *threadService) Delete(thread entity.Threads) {
	service.threadData.Delete(thread)
}

func (service *threadService) List(category string) []entity.Threads {
	return service.threadData.List(category)
}

func (service *threadService) GetPost(tid uint64) entity.Threads {
	return service.threadData.GetPost(tid)
}
