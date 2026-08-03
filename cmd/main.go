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

	// env
	cfg := config.LoadConfig()
	router := gin.Default()

	// database connection
	db, err := database.New(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := respository.NewUserRespository(db)
	service := service.NewUserService(repo, cfg)
	authHandler := handler.NewAuthHandler(service)
	routes.SetupRoutes(router, authHandler)

	// health endpoint
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

	// start server
	router.Run(":" + cfg.Port)
}
