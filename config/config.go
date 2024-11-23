package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	AppPort    string
	JWTKey     string
}

// LoadConfig initializes the application configuration
func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		AppPort:    os.Getenv("APP_PORT"),
		JWTKey:     os.Getenv("JWT_KEY"),
	}
}
