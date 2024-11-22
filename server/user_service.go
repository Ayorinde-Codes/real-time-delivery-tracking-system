package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os/user"

	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/util"
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

func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *user.AuthenticateUserRequest) (*user.AuthenticateUserResponse, error) {
	var id int32
	var hashedPassword string

	// Query the database for the user's details
	err := s.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", req.Email).Scan(&id, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Printf("Error fetching user: %v", err)
		return nil, errors.New("failed to authenticate user")
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate a JWT token
	token, err := util.GenerateJWT(id)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, errors.New("failed to generate token")
	}

	return &user.AuthenticateUserResponse{Token: token}, nil
}
