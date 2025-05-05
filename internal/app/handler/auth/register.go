package auth

import (
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/handler"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()

	data, err := checkRegisterRequest(r)
	if errors.Is(err, handler.ErrUnmarshal) || errors.Is(err, handler.ErrBodyRead) {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if errors.Is(err, handler.ErrValidation) {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var hashPass string
	hashPass, err = hashPassword(data.Password)
	if err != nil {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var UserID int
	UserID, err = h.User.Create(data.Login, hashPass)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		sugar.Infof("User with login %s already exist\n", data.Login)
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.Reward.Create(UserID)
	if err != nil {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var token string
	token, err = CreateToken(data.Login)
	if err != nil {
		sugar.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: token})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func checkRegisterRequest(r *http.Request) (RegisterRequest, error) {
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
