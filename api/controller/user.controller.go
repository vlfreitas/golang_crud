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

// CreateUser		godoc
// @Summary			Create users
// @Description		Save users data in Db.
// @Param			users body models.UserRegister true "Create users"
// @Produce			application/json
// @Tags			users
// @Success			200
// @Router			/users/create [post]
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

// UpdateUserById		godoc
// @Summary			Update users
// @Description		Update users data.
// @Param			users body models.UserUpdate true  "Update user"
// @Param			id path string true "update user by id"
// @Tags			users
// @Produce			application/json
// @Success			200
// @Router			/users/{id} [put]
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
	if result.Error != nil {
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

// FindAllUsers		godoc
// @Summary			Get All users.
// @Description		Return list of users.
// @Tags			users
// @Success			200
// @Router			/users/ [get]
func (u *UserController) ListAll(c *gin.Context) {

	var users []models.User
	results := u.DB.Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// FindUserById 		godoc
// @Summary			Get Single user by id.
// @Param			id path string true "update users by id"
// @Description		Return the users whoes id value match id.
// @Produce			application/json
// @Tags			users
// @Success			200
// @Router			/users/{id} [get]
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

// DeleteUserById		godoc
// @Summary			Delete users
// @Param			id path string true "update users by id"
// @Description		Remove user data by id.
// @Produce			application/json
// @Tags			users
// @Success			200
// @Router			/users/{id} [delete]
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
