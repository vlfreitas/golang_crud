package repository

import (
	"go-crud/infra"
	"go-crud/models"
)

type UserRepository struct {
	db infra.Database
}

func NewUserRepository(db infra.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user models.User) error {

	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.Name = user.Name
	dbUser.Age = user.Age
	dbUser.Password = user.Password
	dbUser.Address = user.Address
	return u.db.DB.Create(&dbUser).Error
}
