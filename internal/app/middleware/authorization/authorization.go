package authorization

import (
	"context"
	"github.com/VladimirSh98/Gophermart.git/internal/app/handler"
	"net/http"
)

func Authorization(handler *handler.Handler) func(h http.Handler) http.Handler {
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
			err = auth.checkUser(handler)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserIDKey, auth.UserID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(logFn)
	}
}

//func createToken(login string) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
//		},
//		UserID: UserCount,
//	})
//	tokenString, err := token.SignedString([]byte(SecretKey))
//	if err != nil {
//		return "", 0, err
//	}
//	return tokenString, nil
//}
//}
