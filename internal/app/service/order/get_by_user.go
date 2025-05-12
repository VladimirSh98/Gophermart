package order

import (
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

func (s *Service) GetByUser(userID int) ([]orderRepo.Order, error) {
	orders, err := s.Repo.GetByUser(userID)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
