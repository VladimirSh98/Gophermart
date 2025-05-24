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
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	sugar := zap.S()
	ctx := r.Context()
	userID := ctx.Value(authorization.UserIDKey).(int)
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

	err = h.checkOrderByID(ctx, orderID, userID)
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
	err = h.Order.Create(ctx, orderID, userID)
	if err != nil {
		sugar.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go h.processingOrder(ctx, orderID, userID)
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) checkOrderByID(ctx context.Context, orderID string, userID int) error {
	var order orderRepo.Order
	var err error
	order, err = h.Order.GetByID(ctx, orderID)
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
