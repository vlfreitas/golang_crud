package models

import (
	"go-crud/utils"

	"gorm.io/gorm"
)

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Age      int    `gorm:"not null" json:"age"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Address  string `gorm:"size:255;not null" json:"address"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = utils.HashPassword(user.Password)
	return nil
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["email"] = user.Email
	resp["name"] = user.Name
	resp["age"] = user.Age
	resp["address"] = user.Address
	return resp
}
