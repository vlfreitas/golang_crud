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

func (u UserRepository) CreateUser(user models.UserRegister) error {

	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.Name = user.Name
	dbUser.Age = user.Age
	dbUser.Password = user.Password
	dbUser.Address = user.Address
	return u.db.DB.Create(&dbUser).Error
}

func (u UserRepository) UpdateUser(user models.User, id int, payload models.UserUpdate) error {

	err := u.db.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return err
	}

	userToUpdate := models.User{
		Name:    payload.Name,
		Age:     payload.Age,
		Email:   payload.Email,
		Address: payload.Address,
	}

	return u.db.DB.Model(&user).Updates(userToUpdate).Error

}

func (u UserRepository) ListAll() (*[]models.User, error) {
	var users []models.User
	err := u.db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil

}

func (u UserRepository) GetById(id int) (*models.User, error) {
	var user models.User

	err := u.db.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) DeleteById(id int) error {
	var user models.User
	return u.db.DB.Delete(&user, "id = ?", id).Error
}
