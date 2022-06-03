package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/yamess/auth/configs"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateToken(userId uuid.UUID, email string) (string, error) {
	expirationTime := time.Now().Add(configs.ExpiredTime * time.Minute)
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"alg":     configs.Algorithm,
		"typ":     "JWT",
		"exp":     expirationTime.Unix(),
		"nbf":     time.Now(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(configs.SecretKey))
	return token, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), configs.HashingCost)
	return string(bytes), err
}

func VerifyPassword(hashPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	return err == nil
}
