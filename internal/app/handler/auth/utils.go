package auth

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateToken(login string) (string, error) {
	exp := time.Hour * time.Duration(config.Conf.TokenExp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
		Login: login,
	})
	tokenString, err := token.SignedString([]byte(config.Conf.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
