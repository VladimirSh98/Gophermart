package order

import (
	"context"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

func (s *Service) GetByUser(ctx context.Context, userID int) ([]orderRepo.Order, error) {
	orders, err := s.Repo.GetByUser(ctx, userID)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
