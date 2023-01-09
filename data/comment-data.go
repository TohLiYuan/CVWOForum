package data

import (
	"App/entity"

	"gorm.io/gorm"
)

func (db *database) SaveComment(comment entity.Comments) {
	db.connection.Create(&comment)
}
func (db *database) DeleteComment(comment entity.Comments) {
	db.connection.Delete(&comment)
}
func (db *database) ListComment(tid uint64) []entity.Comments {
	var comments []entity.Comments
	db.connection.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Omit("email", "password")
	}).Where("Threadid = ?", tid).Find(&comments)
	return comments
}
