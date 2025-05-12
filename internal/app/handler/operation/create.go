package operation

import (
	"encoding/json"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	rewardRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
	"github.com/VladimirSh98/Gophermart.git/internal/app/utils/luhn"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()
	userID := r.Context().Value(authorization.UserIDKey).(int)
	var data CreateRequest
	err := handleRequest(r, &data)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = validateRequest(&data)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	var rewardInfo rewardRepo.Reward
	rewardInfo, err = h.Reward.GetByUser(userID)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var newBalance, newWithdrawn float64
	newBalance, newWithdrawn, err = calculateBalance(rewardInfo, data.Sum)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	err = h.Operation.Create(data.Order, userID, data.Sum)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.Reward.UpdateByUser(userID, newBalance, newWithdrawn)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleRequest(r *http.Request, data *CreateRequest) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return err
	}
	return nil
}

func validateRequest(data *CreateRequest) error {
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		return err
	}
	valid := luhn.IsValid(data.Order)
	if !valid {
		return errors.New("invalid order number")
	}
	return nil
}

func calculateBalance(rewardInfo rewardRepo.Reward, sum float64) (float64, float64, error) {
	if sum > rewardInfo.Balance {
		return 0.0, 0.0, errors.New("not enough balance")
	}
	return rewardInfo.Balance - sum, rewardInfo.Withdrawn + sum, nil
}
