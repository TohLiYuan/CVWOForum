package main

import (
	"App/controller"
	"App/entity"
	"App/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Toh-10030501"
	dbname   = "CVWOForum"
)

var (
	threadService    service.ThreadService       = service.New()
	threadController controller.ThreadController = controller.New(threadService)
	db               *gorm.DB
	err              error
)

func main() {
	server := gin.Default()

	psqlconn := fmt.Sprintf("host=%s user = %s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully connected to database", db)
	}

	db.AutoMigrate(&entity.Users{})
	db.AutoMigrate(&entity.Threads{})
	db.AutoMigrate(&entity.Comments{})

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
