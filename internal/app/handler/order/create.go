package order

import (
	"database/sql"
	"errors"
	"github.com/VladimirSh98/Gophermart.git/internal/app/middleware/authorization"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
	"github.com/VladimirSh98/Gophermart.git/internal/app/utils/luhn"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()
	UserID := r.Context().Value(authorization.UserIDKey).(int)
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		sugar.Errorln("CreateShortURL body read error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	OrderID := string(body)

	valid := luhn.IsValid(OrderID)
	if !valid {
		sugar.Warnln("Create order validation error", OrderID)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = checkOrderByID(h, OrderID, UserID)
	if errors.Is(err, ErrOrderLoadedByAnother) {
		sugar.Warn(err)
		w.WriteHeader(http.StatusConflict)
		return
	} else if errors.Is(err, ErrExistOrder) {
		sugar.Warn(err)
		w.WriteHeader(http.StatusOK)
		return
	} else if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.Order.Create(OrderID, UserID)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go processingOrder(h, OrderID, UserID)
	w.WriteHeader(http.StatusAccepted)
}

func checkOrderByID(h *Handler, OrderID string, UserID int) error {
	var order orderRepo.Order
	var err error
	order, err = h.Order.GetByID(OrderID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		return err
	}
	if order.UserID != UserID {
		return ErrOrderLoadedByAnother
	}
	return ErrExistOrder
}

func processingOrder(h *Handler, OrderID string, UserID int) {
	sugar := zap.S()
	chIn := make(chan string)
	defer close(chIn)
	chDone := make(chan ProcessedResult)
	defer close(chDone)
	go checkStatus(h, chIn, chDone)
	var err error
	for {
		chIn <- OrderID
		result := <-chDone
		switch result.Status {
		case RegisteredStatus:
			time.Sleep(1 * time.Second)
			continue
		case ProcessingStatus:
			err = h.Order.UpdateByID(result.OrderID, result.Status, sql.NullFloat64{Float64: result.Accrual, Valid: true})
			if err != nil {
				sugar.Error(err)
			}
			time.Sleep(1 * time.Second)
			continue
		case InvalidStatus, ProcessedStatus:
			err = h.Order.UpdateByID(result.OrderID, result.Status, sql.NullFloat64{Float64: result.Accrual, Valid: true})
			if err != nil {
				sugar.Error(err)
			}
			err = h.Reward.AccrueReward(UserID, result.Accrual)
			if err != nil {
				sugar.Error(err)
			}
			return
		default:
			err = h.Order.UpdateByID(result.OrderID, InvalidStatus, sql.NullFloat64{Valid: false})
			if err != nil {
				sugar.Error(err)
			}
			return
		}
	}

}

func checkStatus(h *Handler, chIn chan string, chDone chan ProcessedResult) {
	sugar := zap.S()
	for {
		OrderID, ok := <-chIn
		if !ok {
			sugar.Infoln("checkStatus done channel closed")
			return
		}
		result, err := h.Accrual.GetByNumber(OrderID)
		if err != nil {
			sugar.Warn(err)
			chDone <- ProcessedResult{
				OrderID: OrderID,
				Status:  InvalidStatus,
			}
		}
		if result.StatusCode == http.StatusOK {
			chDone <- ProcessedResult{
				OrderID: OrderID,
				Status:  result.Status,
				Accrual: result.Accrual,
			}
		} else if result.StatusCode == http.StatusTooManyRequests {
			time.Sleep(1 * time.Second)
			chDone <- ProcessedResult{
				OrderID: OrderID,
				Status:  RegisteredStatus,
			}
		} else {
			chDone <- ProcessedResult{
				OrderID: OrderID,
				Status:  InvalidStatus,
			}
		}
	}
}
