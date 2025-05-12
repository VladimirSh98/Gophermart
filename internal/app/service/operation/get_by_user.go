package operation

import (
	"context"
	operationRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"
)

func (s *Service) GetByUser(ctx context.Context, userID int) ([]operationRepo.Operation, error) {
	operations, err := s.Repo.GetByUser(ctx, userID)
	if err != nil {
		return operations, err
	}
	return operations, nil
}
