package server

import (
	"context"
	"errors"
	"log"
	"math"

	"github.com/ayorinde-codes/real-time-delivery-tracking/models"
	"github.com/ayorinde-codes/real-time-delivery-tracking/proto/user"
	"github.com/ayorinde-codes/real-time-delivery-tracking/util"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	DB *gorm.DB
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
	if err := s.DB.Create(&newUser).Error; err != nil {
		if isUniqueConstraintError(err) { // Check for unique constraint violation
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return &user.RegisterUserResponse{
		Message: "User registered successfully",
	}, nil
}

func isUniqueConstraintError(err error) bool {
	// for PostgreSql
	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
		return true
	}
	return false
}

// AuthenticateUser handles user authentication and token generation.
func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *user.AuthenticateUserRequest) (*user.AuthenticateUserResponse, error) {
	// Input validation
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Find the user by email
	var existingUser models.User
	if err := s.DB.Where("email = ?", req.Email).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Convert uint to int32 for token generation
	if existingUser.ID > uint(math.MaxInt32) {
		return nil, errors.New("user ID exceeds allowable range for token generation")
	}

	token, err := util.GenerateJWT(int32(existingUser.ID)) // Safely cast to int32
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, errors.New("failed to generate token")
	}

	return &user.AuthenticateUserResponse{
		Token: token,
	}, nil
}
