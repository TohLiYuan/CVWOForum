package service

import (
	"App/data"
	"App/entity"
)

type CommentService interface {
	SaveComment(comment entity.Comments) entity.Comments
	DeleteComment(comment entity.Comments)
	ListComment(tid uint64) []entity.Comments
}

type commentService struct {
	commentData data.Data
}

func NewCommentService(data data.Data) CommentService {
	return &commentService{
		commentData: data,
	}
}

func (service *commentService) SaveComment(comment entity.Comments) entity.Comments {
	service.commentData.SaveComment(comment)
	return comment
}

func (service *commentService) DeleteComment(comment entity.Comments) {
	service.commentData.DeleteComment(comment)
}

func (service *commentService) ListComment(tid uint64) []entity.Comments {
	return service.commentData.ListComment(tid)
}
