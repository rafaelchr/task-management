package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rafaelchr/task-management/config"
)

func GenerateToken(userID int64, role, email string, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTExpired)

	claims := jwt.MapClaims{
		"user_id":   userID,
		"role":      role,
		"email":     email,
		"public_id": publicID,
		"exp":       time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
