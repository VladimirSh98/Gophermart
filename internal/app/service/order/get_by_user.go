package order

import (
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

func (s *Service) GetByUser(UserID int) ([]orderRepo.Order, error) {
	orders, err := s.Repo.GetByUser(UserID)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
