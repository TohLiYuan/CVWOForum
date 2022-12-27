package service

import "App/entity"

type CommentService interface {
	Save(entity.Comments) entity.Comments
	List() []entity.Comments
}

type commentService struct {
	comments []entity.Comments
}

func NewCommentService() CommentService {
	return &commentService{
		comments: []entity.Comments{},
	}
}

func (service *commentService) Save(comment entity.Comments) entity.Comments {
	service.comments = append(service.comments, comment)
	return comment
}

func (service *commentService) List() []entity.Comments {
	return service.comments
}
