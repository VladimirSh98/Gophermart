package order

import (
	"context"
	"database/sql"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

type ServiceInterface interface {
	GetByUser(ctx context.Context, userID int) ([]orderRepo.Order, error)
	GetByID(ctx context.Context, orderID string) (orderRepo.Order, error)
	Create(ctx context.Context, orderID string, userID int) error
	UpdateByID(ctx context.Context, orderID string, status string, value sql.NullFloat64) error
}
