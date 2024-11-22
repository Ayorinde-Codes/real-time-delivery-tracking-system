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
}

// LoadConfig initializes and returns the application configuration
func LoadConfig() *Config {
	// Load environment variables from the `.env` file
	err := godotenv.Load("../config/env") // Update this path if needed
	if err != nil {
		log.Println("Warning: .env file not found, falling back to environment variables")
	}

	return &Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBName:     getEnv("DB_NAME", "real_time_tracking"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		AppPort:    getEnv("APP_PORT", "8080"),
	}
}

// getEnv retrieves environment variables or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
