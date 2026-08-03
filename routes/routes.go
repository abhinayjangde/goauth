package routes

import (
	"github.com/abhinayjangde/goauth/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, auth *handler.AuthHandler) {

	api := router.Group("/api/auth")

	{
		api.POST("/register", auth.Register)
		api.POST("/login", auth.Login)
	}
}
