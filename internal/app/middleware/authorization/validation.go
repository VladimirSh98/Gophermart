package authorization

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	"github.com/golang-jwt/jwt/v4"
)

func (auth *userAuth) validate() error {
	var err error
	auth.token, err = jwt.ParseWithClaims(auth.tokenString, auth, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.SecretKey), nil
	})
	if err != nil {
		return err
	}
	if !auth.token.Valid {
		return ErrNotValidToken
	}
	return nil
}
