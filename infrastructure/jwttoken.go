package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"loantracker/domain" // Replace with the actual path to your domain package

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// GenerateJWT generates a JWT token based on the provided claims
func GenerateJWT(name, id, role string, isaccesstoken bool) (string, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the secret key from the environment variable
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("SECRET_KEY environment variable not set")
	}
	var expie int64

	if isaccesstoken {
		expie = time.Now().Add(time.Hour * 1).Unix()
	} else {
		expie = time.Now().Add(time.Hour * 24 * 10).Unix()
	}

	// Set custom claims
	claims := &domain.Claims{
		Name: name,
		ID:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expie,
		},
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT parses a JWT token and returns the claims
func ParseJWT(tokenString string) (*domain.Claims, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the secret key from the environment variable
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("SECRET_KEY environment variable not set")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims
	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token claims")
	}
}
