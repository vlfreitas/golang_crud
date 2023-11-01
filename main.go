package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	gin.SetMode(gin.TestMode)

	router := server.Group("api/v1")

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://br.pinterest.com/pin/861524603709141306/")
	})

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Up"})
	})
	log.Fatal(server.Run()) // listen and serve on 0.0.0.0:8080
}
