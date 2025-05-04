package auth

import (
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/handler"
	userRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()

	data, err := checkLoginRequest(r)
	if errors.Is(err, handler.ErrUnmarshal) || errors.Is(err, handler.ErrBodyRead) {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if errors.Is(err, handler.ErrValidation) {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user userRepo.User
	user, err = h.User.GetByLogin(data.Login, false)
	if err != nil {
		sugar.Infoln(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func checkLoginRequest(r *http.Request) (RegisterRequest, error) {
	var data RegisterRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return data, handler.ErrBodyRead
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, handler.ErrUnmarshal
	}
	v := validator.New()
	err = v.Struct(data)
	if err != nil {
		return data, handler.ErrValidation
	}
	return data, nil
}
