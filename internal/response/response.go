package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func Success(
	c *gin.Context,
	status int,
	message string,
	data interface{},
) {
	c.JSON(status, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(
	c *gin.Context,
	status int,
	message string,
	errors map[string]string,
) {
	c.JSON(status, ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Success: false,
		Message: "Internal server error",
	})
}
