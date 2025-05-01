package authorization

import "github.com/golang-jwt/jwt/v4"

type contextKey string

const UserIDKey contextKey = "userID"

type userAuth struct {
	jwt.RegisteredClaims
	tokenString string
	token       *jwt.Token
	login       string
	password    string
	userID      int
}
