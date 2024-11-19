package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "bonnoylinda"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token differs from signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("error parsing jwt token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("token differs from signing method")
	}

	// email, _ := claims["email"].(string)
	userID := int64(claims["userId"].(float64))
	return userID, nil
}