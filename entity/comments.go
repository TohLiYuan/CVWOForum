package entity

type Comments struct {
	ID       int64  `gorm:"primarykey;default:nextval('comments_id_seq'::regclass)"`
	Text     string `json:"text" binding:"required,max=2000" gorm:"not null;size:2000"`
	Epoch    int64  `gorm:"autoCreateTime"`
	Userid   int64
	Users    Users `gorm:"foreignKey:Userid"`
	Threadid int64
	Threads  Threads `gorm:"foreignKey:Threadid"`
}
