package authorization

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/handler"
)

func (auth *userAuth) checkUser(handler *handler.Handler) error {
	user, err := handler.User.GetByLogin(auth.Login, false)
	if err != nil {
		return err
	}
	auth.UserID = user.ID
	return nil
}
