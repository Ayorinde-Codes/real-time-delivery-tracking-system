package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // Pgx driver for SQL
)

func Connect() (*sql.DB, error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}
