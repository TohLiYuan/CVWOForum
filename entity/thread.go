package entity

type Threads struct {
	ID      int64  `gorm:"primarykey;default:nextval('threads_id_seq'::regclass)"`
	Title   string `json:"title" binding:"required,min=2,max=256" gorm:"not null;size:256"`
	Content string `json:"content"`
	URL     string `json:"url" binding:"required,url" gorm:"not null"`
	Userid  int64
	Users   Users `gorm:"foreignKey:Userid"`
	Epoch   int64 `gorm:"autoCreateTime"`
}
