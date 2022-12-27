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
	password = "Toh-10030501"
	dbname   = "CVWOForum"
)

type Data interface {
	CloseDB()
}

type ThreadData interface {
	Save(thread entity.Threads)
	Update(thread entity.Threads)
	Delete(thread entity.Threads)
	List() []entity.Threads
}

type database struct {
	connection *gorm.DB
}

func NewData() Data {
	psqlconn := fmt.Sprintf("host=%s user = %s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Users{}, &entity.Threads{}, &entity.Comments{})

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
