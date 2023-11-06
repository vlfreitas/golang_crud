package main

import (
	"go-crud/api/controller"
	"go-crud/api/routes"
	"go-crud/infra"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Env
	infra.InitEnv()

	// Init DB
	db := infra.InitDb()

	// Start Controller and Route User
	userController := controller.NewUserController(db.DB)
	userRouteController := routes.NewUserRouteController(userController)

	// Setup Server
	server := gin.Default()
	server.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "App Running!"})
	})

	// Setup Router Api
	router := server.Group("/api/v1")
	userRouteController.Setup(router)

	// Migrate DB
	db.DB.AutoMigrate(&models.User{})
	log.Print("Database migrate done!")

	// Run Server
	log.Fatal(server.Run())
}
