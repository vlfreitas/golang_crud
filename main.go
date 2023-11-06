package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-crud/api/controller"
	"go-crud/api/routes"
	_ "go-crud/docs"
	"go-crud/infra"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title 	Users Service API
// @version	1.0
// @description A User service API in Go using Gin framework
// @host 	localhost:8080
// @BasePath /api/v1
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
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "App Running!"})
	})

	// Setup Router Api
	baseRouter := server.Group("/api/v1")
	usersRouter := baseRouter.Group("/users")
	userRouteController.Setup(usersRouter)

	// Migrate DB
	db.DB.AutoMigrate(&models.User{})
	log.Print("Database migrate done!")

	// Run Server
	log.Fatal(server.Run())
}
