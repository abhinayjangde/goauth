package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl    string
	JwtSecret      string
	Port           string
	JWTExpireHours string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using system environment")
	}

	return &Config{
		DatabaseUrl:    os.Getenv("DATABASE_URL"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		Port:           os.Getenv("PORT"),
		JWTExpireHours: os.Getenv("JWT_EXPIRE_HOURS"),
	}
}
