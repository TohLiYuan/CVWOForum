package data

import (
	"App/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = ""
)

type Data interface {
	CloseDB()

	Save(thread entity.Threads)
	Update(thread entity.Threads)
	Delete(thread entity.Threads)
	List(cat string) []entity.Threads
	GetPost(tid uint64) entity.Threads

	SaveComment(comment entity.Comments)
	DeleteComment(comment entity.Comments)
	ListComment(tid uint64) []entity.Comments

	SaveUser(user entity.Users)
}

var DB *gorm.DB

type database struct {
	connection *gorm.DB
}

func NewData() Data {
	psqlconn := fmt.Sprintf("host=%s user = %s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	var err error

	DB, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&entity.Users{}, &entity.Threads{}, &entity.Comments{})

	return &database{
		connection: DB,
	}
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
