package order

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func (h *Handler) processingOrder(ctx context.Context, orderID string, userID int) {
	sugar := zap.S()
	var err error
	for {
		result := h.checkStatus(ctx, orderID)
		switch result.Status {
		case RegisteredStatus:
			time.Sleep(1 * time.Second)
			continue
		case ProcessingStatus:
			err = h.Order.UpdateByID(ctx, result.OrderID, result.Status, sql.NullFloat64{Float64: result.Accrual, Valid: true})
			if err != nil {
				sugar.Error(err)
			}
			time.Sleep(1 * time.Second)
			continue
		case InvalidStatus, ProcessedStatus:
			err = h.Order.UpdateByID(ctx, result.OrderID, result.Status, sql.NullFloat64{Float64: result.Accrual, Valid: true})
			if err != nil {
				sugar.Error(err)
			}
			err = h.Reward.AccrueReward(ctx, userID, result.Accrual)
			if err != nil {
				sugar.Error(err)
			}
			return
		default:
			err = h.Order.UpdateByID(ctx, result.OrderID, InvalidStatus, sql.NullFloat64{Valid: false})
			if err != nil {
				sugar.Error(err)
			}
			return
		}
	}

}

func (h *Handler) checkStatus(ctx context.Context, orderID string) ProcessedResult {
	for {
		result, err := h.Accrual.GetByNumber(ctx, orderID)
		if err != nil {
			continue
		}
		if result.StatusCode == http.StatusOK {
			return ProcessedResult{
				OrderID: orderID,
				Status:  result.Status,
				Accrual: result.Accrual,
			}
		} else if result.StatusCode == http.StatusTooManyRequests {
			time.Sleep(1 * time.Second)
			return ProcessedResult{
				OrderID: orderID,
				Status:  RegisteredStatus,
			}
		} else {
			return ProcessedResult{
				OrderID: orderID,
				Status:  InvalidStatus,
			}
		}
	}
}
