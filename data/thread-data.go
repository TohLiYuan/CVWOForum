package data

import (
	"App/entity"

	"gorm.io/gorm"
)

func (db *database) Save(thread entity.Threads) {
	db.connection.Create(&thread)
}

func (db *database) Update(thread entity.Threads) {
	db.connection.Save(&thread)
}
func (db *database) Delete(thread entity.Threads) {
	db.connection.Delete(&thread)
}

func (db *database) List(cat string) []entity.Threads {
	var threads []entity.Threads
	if cat == "all" {
		db.connection.Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("email", "password")
		}).Find(&threads)
	} else {
		db.connection.Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("email", "password")
		}).Where("category = ?", cat).Find(&threads)
	}

	return threads
}

func (db *database) GetPost(tid uint64) entity.Threads {
	var t entity.Threads
	db.connection.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Omit("email", "password")
	}).Where("ID = ?", tid).Find(&t)

	return t
}
