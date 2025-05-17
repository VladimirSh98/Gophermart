package operation

import (
	"context"
	"github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
)

type ServiceInterface interface {
	GetByUser(ctx context.Context, userID int) ([]operation.Operation, error)
	Create(ctx context.Context, orderID string, userID int, value float64) error
}
