package main

import (
	"net/http"

	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	router.Run(":" + cfg.Port)
}
