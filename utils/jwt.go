package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId int64) (string, error) {

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId" : email,
		"email": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}