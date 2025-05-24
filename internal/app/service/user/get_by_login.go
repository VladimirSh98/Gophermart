package user

import (
	"context"
	userRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
)

func (s *Service) GetByLogin(ctx context.Context, login string, archived bool) (userRepo.User, error) {
	user, err := s.Repo.GetUserByLogin(ctx, login, archived)
	if err != nil {
		return user, err
	}
	return user, nil
}
