package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/abhinayjangde/goauth/internal/database"
	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/abhinayjangde/goauth/internal/respository"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	db, err := database.New(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := respository.NewUserRespository(db)

	router := gin.Default()

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

	router.POST("/api/users", func(c *gin.Context) {
		user := &model.User{
			Name:     "Abhi",
			Email:    "abhi@gmail.com",
			Password: "abhi1234",
		}

		err := repo.Create(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user created successfully",
		})

	})

	router.Run(":" + cfg.Port)
}
