package routes

import (
	"github.com/abhinayjangde/goauth/internal/config"
	"github.com/abhinayjangde/goauth/internal/handler"
	"github.com/abhinayjangde/goauth/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, auth *handler.AuthHandler) {

	api := router.Group("/api/auth")

	{
		api.POST("/register", auth.Register)
		api.POST("/login", auth.Login)
	}

	protected := router.Group("/api")
	protected.Use(
		middleware.AuthMiddleware(config.LoadConfig().JwtSecret),
	)
	protected.GET(
		"/profile",
		auth.Profile,
	)
}
