package authorization

import (
	"context"
	authHandler "github.com/VladimirSh98/Gophermart.git/internal/app/handler/auth"
	"net/http"
)

func Authorization(handler *authHandler.Handler) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		logFn := func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("Authorization")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			auth := &userAuth{tokenString: cookie.Value}
			err = auth.validate()
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := r.Context()
			err = auth.checkUser(ctx, handler)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx = context.WithValue(ctx, UserIDKey, auth.UserID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(logFn)
	}
}
