package main

import (
	"go-crud/infra"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "App Running!"})
	})

	// Load Env
	infra.InitEnv()

	// Init DB and Migrate Models
	db := infra.InitDb()
	db.DB.AutoMigrate(&models.User{})

	log.Fatal(server.Run()) // listen and serve on 0.0.0.0:8080
}
