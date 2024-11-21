package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os/user"

	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
	DB *sql.DB
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Failed to hash password")
	}

	_, err := s.DB.Exec("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", req.Name, req.Email, string(hashedPassword), req.Role)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, errors.New("Failed to register user")
	}
	return &user.RegisterUserResponse{Message: "User registered successfully"}, nil
}
