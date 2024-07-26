package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId int64) (string, error) {

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId" : userId,
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (int64 , error) {
	// Read the secret from the .env file
	secret := os.Getenv("JWT_SECRET")

	// Parse the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Return the secret used to sign the token
		return []byte(secret), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	// Check if parsed token is valid
	if !parsedToken.Valid {
		return 0, errors.New("token is not valid")
	}

	// Check if the parsed token is a map claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("could not parse claims")
	}

	// Check if the token has expired
	exp := claims["exp"].(float64)

	if time.Now().Unix() > int64(exp) {
		return 0,errors.New("token has expired")
	}

	// Get the userId from the token
	userId := int64(claims["userId"].(float64))

	return userId, nil
}