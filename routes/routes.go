package routes

import (
	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/abhinayjangde/goauth/internal/handler"
	"github.com/abhinayjangde/goauth/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, auth *handler.AuthHandler, cfg *config.Config) {

	api := router.Group("/api")

	authGroup := api.Group("/auth")

	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)
	authGroup.POST("/refresh", auth.Refresh)

	protected := api.Group("/")
	protected.Use(
		middleware.AuthMiddleware(cfg.JwtSecret),
	)

	protected.GET("/profile", auth.Profile)
}
