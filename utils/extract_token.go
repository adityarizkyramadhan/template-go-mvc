package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetExpiredToken(tokenString string) (time.Duration, error) {
	// Parse the JWT token
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, err
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("could not parse claims")
	}

	// Get the expiration time
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, fmt.Errorf("could not find exp claim")
	}

	// Calculate the time until expiration
	expirationTime := time.Unix(int64(exp), 0)
	duration := time.Until(expirationTime)

	return duration, nil
}
