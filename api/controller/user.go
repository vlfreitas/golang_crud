package controller

import (
	"go-crud/api/service"
	"go-crud/models"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

// NewUserController : NewUserController
func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

// CreateUser ->  calls CreateUser services for validated user
func (u *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	hashPassword := utils.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK)
}
