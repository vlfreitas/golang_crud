package models

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Age      int    `gorm:"not null" json:"age"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Address  string `gorm:"size:255;not null" json:"address"`
}
