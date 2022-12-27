package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Userid   int64  `gorm:"primarykey;default:nextval('users_userid_seq'::regclass)"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}
