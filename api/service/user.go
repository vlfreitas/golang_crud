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

func (u UserService) CreateUser(user models.User) error {
	return u.repo.CreateUser(user)
}
