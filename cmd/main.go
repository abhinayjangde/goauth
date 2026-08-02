package main

import (
	"fmt"
	"net/http"

	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	fmt.Println("database url", cfg.DatabaseUrl)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	router.Run(":8080")
}
