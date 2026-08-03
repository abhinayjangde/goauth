package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64, email string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}
