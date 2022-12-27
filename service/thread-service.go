package service

import "App/entity"

type ThreadService interface {
	Save(entity.Threads) entity.Threads
	List() []entity.Threads
}

type threadService struct {
	threads []entity.Threads
}

func New() ThreadService {
	return &threadService{
		threads: []entity.Threads{},
	}
}

func (service *threadService) Save(thread entity.Threads) entity.Threads {
	service.threads = append(service.threads, thread)
	return thread
}

func (service *threadService) List() []entity.Threads {
	return service.threads
}
