package auth

import (
	"encoding/json"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetByUser(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()
	UserID := r.Context().Value(authorization.UserIDKey).(int)
	res, err := h.Reward.GetByUser(UserID)
	if err != nil {
		sugar.Errorln("GetByUser error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := GetByUserResponse{
		Current:   res.Balance,
		Withdrawn: res.Withdrawn,
	}
	var response []byte
	response, err = json.Marshal(result)
	if err != nil {
		sugar.Warnln("GetByUser json marshall error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		sugar.Errorln("GetByUser response error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
