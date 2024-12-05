package server_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/server"
	_ "github.com/jackc/pgx/v4/stdlib" // PostgreSQL driver
)

func TestRegisterUser(t *testing.T) {
	// Setup
	db, err := sql.Open("pgx", "your_test_db_connection_string")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	s := server.UserServiceServer{DB: db}

	// Test case
	req := &user.RegisterUserRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}

	res, err := s.RegisterUser(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	if res.Message != "User registered successfully!" {
		t.Errorf("Unexpected response message: %v", res.Message)
	}
}
