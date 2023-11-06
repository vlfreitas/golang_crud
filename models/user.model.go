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
	Password string `gorm:"size:100;not null;" json:"-"`
	Address  string `gorm:"size:255;not null" json:"address"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = utils.HashPassword(user.Password)
	return nil
}

type UserRegister struct {
	Name     string `form:"name" gorm:"name" json:"name" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Address  string `form:"address" json:"address" binding:"required"`
}

type UserUpdate struct {
	Name    string `form:"name"json:"name"`
	Age     int    `form:"age" json:"age"`
	Email   string `form:"email" json:"email"`
	Address string `form:"address" json:"address" `
}
