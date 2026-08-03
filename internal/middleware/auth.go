package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})

			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})

			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, errors.New("unexpected signing method")
				}

				return []byte(secret), nil

			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})

			c.Abort()

			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"]
		email := claims["email"]

		// store data in context
		c.Set("user_id", userId)
		c.Set("email", email)
		c.Next()
	}
}
