package controller

import (
	"App/data"
	"App/entity"
	"App/middleware"
	"App/service"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	SaveComment(ctx *gin.Context) error
	DeleteComment(ctx *gin.Context) error
	ListComment(ctx *gin.Context) ([]entity.Comments, error)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(service service.CommentService) CommentController {
	return &commentController{
		commentService: service,
	}
}

func (c *commentController) SaveComment(ctx *gin.Context) error {
	var com entity.Comments
	var t entity.Threads
	var u entity.Users
	var a entity.Users

	// Find user details for logged Uid
	data.DB.Where("user_id = ?", middleware.Uid).First(&u)
	com.Users = u

	tid, err := strconv.ParseUint(ctx.Param("tid"), 0, 0)
	if err != nil {
		return err
	}

	com.Threadid = tid

	// Autofill required fields
	data.DB.Where("ID = ?", tid).First(&t)
	data.DB.Where("user_id = ?", t.Uid).First(&a)
	t.Users = a
	com.Threads = t

	err = ctx.ShouldBindJSON(&com)
	if err != nil {
		return err
	}
	c.commentService.SaveComment(com)

	return nil
}

func (c *commentController) DeleteComment(ctx *gin.Context) error {
	var com entity.Comments

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	data.DB.Where("ID = ?", id).First(&com)

	// Validate user performing delete action
	if com.Userid != middleware.Uid {
		return errors.New("Unauthorised access")
	}

	c.commentService.DeleteComment(com)

	return nil
}

func (c *commentController) ListComment(ctx *gin.Context) ([]entity.Comments, error) {
	tid, err := strconv.ParseUint(ctx.Param("tid"), 0, 0)
	var emp []entity.Comments
	if err != nil {
		return emp, err
	}

	return c.commentService.ListComment(tid), nil
}
