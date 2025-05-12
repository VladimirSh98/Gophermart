package operation

import (
	operationRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
)

func (s *Service) GetByUser(userID int) ([]operationRepo.Operation, error) {
	operations, err := s.Repo.GetByUser(userID)
	if err != nil {
		return operations, err
	}
	return operations, nil
}
