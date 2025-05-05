package operation

import (
	operationRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
)

func (s *Service) GetByUser(UserID int) ([]operationRepo.Operation, error) {
	operations, err := s.Repo.GetByUser(UserID)
	if err != nil {
		return operations, err
	}
	return operations, nil
}
