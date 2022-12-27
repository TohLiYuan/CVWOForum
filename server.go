package main

import (
	"App/controller"
	"App/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	threadService    service.ThreadService       = service.NewThreadService()
	threadController controller.ThreadController = controller.New(threadService)
)

func main() {
	server := gin.Default()

	server.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, threadController.List())
	})

	server.POST("/post", func(ctx *gin.Context) {
		err := threadController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Thread posted!"})
		}

	})

	server.Run()
}
