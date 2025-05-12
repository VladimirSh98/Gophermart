package accrual

import (
	"context"
	"github.com/VladimirSh98/Gophermart.git/internal/app/client/accrual"
)

func (s Service) GetByNumber(ctx context.Context, number string) (*accrual.Calculations, error) {
	res, err := s.Client.GetByNumber(ctx, number)
	if err != nil {
		return res, err
	}
	return res, nil
}
