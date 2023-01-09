package main

import (
	"App/controller"
	"App/data"
	"App/middleware"
	"App/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	fullData data.Data = data.NewData()

	threadService  service.ThreadService  = service.NewThreadService(fullData)
	userService    service.UserService    = service.NewUserService(fullData)
	commentService service.CommentService = service.NewCommentService(fullData)
	jWtService     service.JWTService     = service.NewJWTService()

	threadController  controller.ThreadController  = controller.NewThreadController(threadService)
	userController    controller.UserController    = controller.NewUserController(userService, jWtService)
	commentController controller.CommentController = controller.NewCommentController(commentService)
)

func main() {
	server := gin.Default()

	server.POST("/register", func(ctx *gin.Context) {
		err := userController.Register(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User registered!"})
		}
	})

	server.POST("/login", func(ctx *gin.Context) {
		token, err := userController.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	})

	server.GET("/list/:category", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, threadController.List(ctx))
	})

	server.GET("/:tid", func(ctx *gin.Context) {
		t, err := threadController.GetPost(ctx)
		c, err := commentController.ListComment(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"thread": t})
			ctx.JSON(http.StatusOK, gin.H{"comments": c})
		}
	})

	api := server.Group("/api", middleware.AuthoriseJWT())
	{

		api.POST("/post", func(ctx *gin.Context) {
			err := threadController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Thread posted!"})
			}

		})

		api.PUT("/t/:id", func(ctx *gin.Context) {
			err := threadController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Thread posted"})
			}
		})

		api.DELETE("/t/:id", func(ctx *gin.Context) {
			err := threadController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Thread deleted"})
			}
		})

		api.POST("/comment/:tid", func(ctx *gin.Context) {
			err := commentController.SaveComment(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Comment posted!"})
			}
		})

		api.DELETE("/comment/:id", func(ctx *gin.Context) {
			err := commentController.DeleteComment(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
			}
		})
	}

	server.Run()
}
