package operation

import (
	"encoding/json"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	operationRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetByUser(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()
	ctx := r.Context()
	userID := ctx.Value(authorization.UserIDKey).(int)
	res, err := h.Operation.GetByUser(ctx, userID)
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

func createResponse(res []operationRepo.Operation) []GetByUserResponse {
	var results []GetByUserResponse
	for _, record := range res {
		results = append(results, GetByUserResponse{
			Order:       record.ID,
			Sum:         record.Value,
			ProcessedAt: record.CreatedAt,
		})
	}
	return results
}
