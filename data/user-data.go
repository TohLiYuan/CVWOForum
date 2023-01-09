package data

import (
	"App/entity"
)

func (db *database) SaveUser(user entity.Users) {

	db.connection.Create(&user)

}
