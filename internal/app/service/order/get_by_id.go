package order

import (
	"context"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

func (s *Service) GetByID(ctx context.Context, orderID string) (orderRepo.Order, error) {
	order, err := s.Repo.GetByID(ctx, orderID)
	if err != nil {
		return order, err
	}
	return order, nil
}
