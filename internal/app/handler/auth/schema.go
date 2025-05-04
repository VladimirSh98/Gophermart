package auth

import "github.com/golang-jwt/jwt/v4"

type RegisterRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type tokenData struct {
	jwt.RegisteredClaims
	Login string
}
