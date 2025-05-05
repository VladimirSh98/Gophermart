package authorization

import (
	authHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/auth"
)

func (auth *userAuth) checkUser(handler *authHandler.Handler) error {
	user, err := handler.User.GetByLogin(auth.Login, false)
	if err != nil {
		return err
	}
	auth.UserID = user.ID
	return nil
}
