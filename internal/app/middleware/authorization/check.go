package authorization

import (
	"context"
	authHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/auth"
)

func (auth *userAuth) checkUser(ctx context.Context, handler *authHandler.Handler) error {
	user, err := handler.User.GetByLogin(ctx, auth.Login, false)
	if err != nil {
		return err
	}
	auth.UserID = user.ID
	return nil
}
