package controller

import (
	"go-crud/api/service"
	"go-crud/models"
	"go-crud/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u *UserController) Create(c *gin.Context) {
	var payload models.UserRegister

	log.Print("Checking Payload")
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashPassword := utils.HashPassword(payload.Password)
	payload.Password = hashPassword

	log.Print("Creating User")
	err := u.service.CreateUser(payload)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	log.Print("User created")

	c.JSON(http.StatusOK, gin.H{})
}

func (u *UserController) Update(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var payload models.UserUpdate

	log.Print("Checking Payload")
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var updatedUser models.User
	result := u.service.UpdateUser(updatedUser, id, payload)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that id exists"})
		return
	}

}

func (u *UserController) ListAll(c *gin.Context) {
	users, err := u.service.ListALl()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}

func (u *UserController) GetById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := u.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that id exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (u *UserController) DeleteById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = u.service.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that id exists"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
