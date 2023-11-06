package main

import (
	"go-crud/api/controller"
	"go-crud/api/repository"
	"go-crud/api/routes"
	"go-crud/api/service"
	"go-crud/infra"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "App Running!"})
	})

	// Load Env
	infra.InitEnv()

	// Init DB
	db := infra.InitDb()

	// Setup User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	// Migrate
	db.DB.AutoMigrate(&models.User{})
	log.Print("Database migrate done!")

	log.Fatal(router.Run()) // listen and serve on 0.0.0.0:8080
}
