package handler

import (
	"net/http"

	"github.com/abhinayjangde/goauth/internal/model"
	"github.com/abhinayjangde/goauth/internal/response"
	"github.com/abhinayjangde/goauth/internal/service"
	"github.com/abhinayjangde/goauth/internal/validator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(s *service.UserService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req model.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"validation failed",
			validator.FormatErrors(err),
		)
		return
	}

	err := h.service.Register(&req)

	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"validation failed",
			validator.FormatErrors(err),
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"User registered successfully",
		nil,
	)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	res, err := h.service.Login(&req)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"access_token":  res.AccessToken,
		"refresh_token": res.RefreshToken,
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {

	userId, _ := c.Get("user_id")
	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
		"email":   email,
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {

	var req model.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	res, err := h.service.Refresh(&req)

	if err != nil {

		c.JSON(http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			})

		return
	}

	c.JSON(http.StatusOK, res)
}
