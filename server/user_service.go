package server

import (
	"context"
	"errors"
	"log"

	"github.com/ayorinde-codes/real-time-delivery-tracking/models"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	db *gorm.DB
	user.UnimplementedUserServiceServer
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	// Input validation
	if req.Name == "" || req.Email == "" || req.Password == "" || req.Role == "" {
		return nil, errors.New("all fields (name, email, password, role) are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create a new user model
	newUser := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: string(hashedPassword),
	}

	// Save the user in the database
	if err := s.db.Create(&newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrUniqueConstraint) {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return &user.RegisterUserResponse{
		Message: "User registered successfully",
	}, nil
}

// AuthenticateUser handles user authentication and token generation.
func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *user.AuthenticateUserRequest) (*user.AuthenticateUserResponse, error) {
	// Input validation
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Find the user by email
	var existingUser models.User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := util.GenerateJWT(existingUser.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, errors.New("failed to generate token")
	}

	return &user.AuthenticateUserResponse{
		Token: token,
	}, nil
}
