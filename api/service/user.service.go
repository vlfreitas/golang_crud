package service

import (
	"go-crud/api/repository"
	"go-crud/models"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

func (u UserService) UpdateUser(user models.User, id int, payload models.UserUpdate) error {
	return u.repo.UpdateUser(user, id, payload)
}

func (u UserService) ListALl() (*[]models.User, error) {
	return u.repo.ListAll()
}

func (u UserService) GetById(id int) (*models.User, error) {
	return u.repo.GetById(id)
}

func (u UserService) DeleteById(id int) error {
	return u.repo.DeleteById(id)
}
