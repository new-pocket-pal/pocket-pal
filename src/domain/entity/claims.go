package entity

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID string `json:"user_id"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
