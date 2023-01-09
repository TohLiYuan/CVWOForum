package entity

type Users struct {
	UserID    uint64 `json:"-" gorm:"primarykey;default:nextval('users_userid_seq'::regclass)"`
	Username  string `json:"username" gorm:"not null;default:(concat('user',Userid))"`
	Email     string `json:"email" binding:"required,email" gorm:"not null;UNIQUE"`
	Password  string `json:"password" binding:"required" gorm:"not null"`
	CreatedAt uint64 `gorm:"autoCreateTime"`
}

type Threads struct {
	ID       uint64 `gorm:"primarykey;default:nextval('threads_id_seq'::regclass)" json:"id"`
	Title    string `json:"title" binding:"required,min=2,max=256" gorm:"not null;size:256"`
	Content  string `json:"content"`
	URL      string `json:"url" binding:"required,url" gorm:"not null;default:(concat('https://cvwoforums.com/thread/',ID))"`
	Category string `json:"category" binding:"required" validate:"check-category"`
	Epoch    uint64 `json:"created_at" gorm:"autoCreateTime"`
	Uid      uint64
	Users    Users `json:"users" binding:"required" gorm:"foreignKey:Uid;references:UserID"`
}

type Comments struct {
	ID       uint64 `gorm:"primarykey;default:nextval('comments_id_seq'::regclass)"`
	Text     string `json:"text" binding:"required,max=2000" gorm:"not null;size:2000"`
	Epoch    int64  `gorm:"autoCreateTime"`
	Userid   uint64
	Users    Users `gorm:"foreignKey:Userid;references:UserID"`
	Threadid uint64
	Threads  Threads `json:"-" gorm:"foreignKey:Threadid;references:ID"`
}
