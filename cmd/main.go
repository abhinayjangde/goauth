package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/abhinayjangde/goauth/internal/database"
	"github.com/abhinayjangde/goauth/internal/handler"
	"github.com/abhinayjangde/goauth/internal/respository"
	"github.com/abhinayjangde/goauth/internal/service"
	"github.com/abhinayjangde/goauth/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	db, err := database.New(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	repo := respository.NewUserRespository(db)
	service := service.NewUserService(repo)
	authHandler := handler.NewAuthHandler(service)
	routes.SetupRoutes(router, authHandler)

	router.GET("/health", func(c *gin.Context) {
		version, err := database.CheckConnection(db)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "Error while getting database version",
			})
		}

		version = strings.Join(strings.Split(version, " ")[:2], " ")

		c.JSON(http.StatusOK, gin.H{
			"message":         "Server is running",
			"postgresVersion": version,
		})
	})

	router.Run(":" + cfg.Port)
}
