package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJWT(userId int32) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   string(rune(userId)), // Convert userId to a string
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
