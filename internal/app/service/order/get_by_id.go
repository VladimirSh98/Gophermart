package order

import (
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

func (s *Service) GetByID(orderID string) (orderRepo.Order, error) {
	order, err := s.Repo.GetByID(orderID)
	if err != nil {
		return order, err
	}
	return order, nil
}
