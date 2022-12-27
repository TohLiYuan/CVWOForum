package entity

import "gorm.io/gorm"

type Threads struct {
	ID      int64  `gorm:"primarykey;auto_increment" json:"id"`
	Title   string `json:"title" binding:"required,min=2,max=256" gorm:"not null;size:256"`
	Content string `json:"content"`
	URL     string `json:"url" binding:"required,url" gorm:"not null"`
	Userid  int64
	Users   SecureUsers `gorm:"foreignKey:Userid"`
	Epoch   int64       `gorm:"autoCreateTime"`
}

type Users struct {
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type SecureUsers struct {
	gorm.Model

	Userid   int64  `gorm:"primarykey;auto_increment" json:"uid"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Comments struct {
	ID       int64  `gorm:"primarykey;auto_increment" json:"id"`
	Text     string `json:"text" binding:"required,max=2000" gorm:"not null;size:2000"`
	Epoch    int64  `gorm:"autoCreateTime"`
	Userid   int64
	Users    SecureUsers `gorm:"foreignKey:Userid"`
	Threadid int64
	Threads  Threads `gorm:"foreignKey:Threadid"`
}
