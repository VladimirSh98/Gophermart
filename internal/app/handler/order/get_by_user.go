package auth

import (
	"encoding/json"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetByUser(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()
	UserID := r.Context().Value(authorization.UserIDKey).(int)
	res, err := h.Order.GetByUser(UserID)
	if err != nil {
		sugar.Errorln("GetByUser error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(res) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	results := createResponse(res)
	var response []byte
	response, err = json.Marshal(results)
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

func createResponse(res []orderRepo.Order) []GetByUserResponse {
	var results []GetByUserResponse
	for _, record := range res {
		if record.Value.Valid {
			results = append(results, GetByUserResponse{
				Number:     record.ID,
				UploadedAt: record.UploadedAt,
				Status:     record.Status,
				Accrual:    record.Value.Float64,
			})
		} else {
			results = append(results, GetByUserResponse{
				Number:     record.ID,
				UploadedAt: record.UploadedAt,
				Status:     record.Status,
			})
		}
	}
	return results
}
