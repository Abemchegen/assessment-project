package domain

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
