package controller

import (
	"go-crud/models"
	"go-crud/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
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

	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Age:      payload.Age,
		Password: payload.Password,
		Address:  payload.Address,
	}

	log.Print("Creating User")
	result := u.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User with that email already exists"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"message": result.Error.Error()})
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
	result := u.DB.First(&updatedUser, "id = ?", id)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user with that id exists"})
		return
	}

	userToUpdate := models.User{
		Name:    payload.Name,
		Age:     payload.Age,
		Email:   payload.Email,
		Address: payload.Address,
	}

	u.DB.Model(&updatedUser).Updates(userToUpdate)

	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func (u *UserController) ListAll(c *gin.Context) {

	var users []models.User
	results := u.DB.Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u *UserController) GetById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	result := u.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user with that id exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserController) DeleteById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	result := u.DB.Delete(&user, "id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user with that id exists"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
