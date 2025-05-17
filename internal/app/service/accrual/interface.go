package accrual

import (
	"context"
	"github.com/VladimirSh98/Gophermart.git/internal/app/client/accrual"
)

type ServiceInterface interface {
	GetByNumber(ctx context.Context, number string) (*accrual.Calculations, error)
}
