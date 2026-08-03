package routes

import (
	"github.com/abhinayjangde/goauth/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, auth *handler.AuthHandler) {

	api := router.Group("/api/")

	{
		api.POST("/users", auth.Register)
	}
}
