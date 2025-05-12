package order

import (
	"context"
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
	userID := r.Context().Value(authorization.UserIDKey).(int)
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		sugar.Errorln("CreateShortURL body read error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	orderID := string(body)

	valid := luhn.IsValid(orderID)
	if !valid {
		sugar.Warnln("Create order validation error", orderID)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = checkOrderByID(h, orderID, userID)
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
	err = h.Order.Create(orderID, userID)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go processingOrder(h, orderID, userID)
	w.WriteHeader(http.StatusAccepted)
}

func checkOrderByID(h *Handler, orderID string, userID int) error {
	var order orderRepo.Order
	var err error
	order, err = h.Order.GetByID(orderID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		return err
	}
	if order.UserID != userID {
		return ErrOrderLoadedByAnother
	}
	return ErrExistOrder
}

func processingOrder(h *Handler, orderID string, userID int) {
	sugar := zap.S()
	chIn := make(chan string)
	defer close(chIn)
	chDone := make(chan ProcessedResult)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go checkStatus(ctx, h, chIn, chDone)
	var err error
	for {
		chIn <- orderID
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
			err = h.Reward.AccrueReward(userID, result.Accrual)
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

func checkStatus(ctx context.Context, h *Handler, chIn chan string, chDone chan ProcessedResult) {
	sugar := zap.S()
	defer close(chDone)
	for {
		select {
		case <-ctx.Done():
			sugar.Infoln("checkStatus context cancelled")
			return
		case orderID, ok := <-chIn:
			if !ok {
				sugar.Infoln("checkStatus input channel closed")
				return
			}
			result, err := h.Accrual.GetByNumber(orderID)
			if err != nil {
				continue
			}
			if result.StatusCode == http.StatusOK {
				chDone <- ProcessedResult{
					OrderID: orderID,
					Status:  result.Status,
					Accrual: result.Accrual,
				}
			} else if result.StatusCode == http.StatusTooManyRequests {
				time.Sleep(1 * time.Second)
				chDone <- ProcessedResult{
					OrderID: orderID,
					Status:  RegisteredStatus,
				}
			} else {
				chDone <- ProcessedResult{
					OrderID: orderID,
					Status:  InvalidStatus,
				}
			}
		}
	}
}
