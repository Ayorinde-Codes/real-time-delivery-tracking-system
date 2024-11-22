package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ayorinde-codes/real-time-delivery-tracking/config"
	_ "github.com/jackc/pgx/v5/stdlib" // Import pgx driver
)

// Connect initializes and returns a database connection
func Connect(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	// Open a database connection
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Database connected successfully!")
	return db, nil
}
