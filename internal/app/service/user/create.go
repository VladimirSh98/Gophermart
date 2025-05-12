package user

import "context"

func (s *Service) Create(ctx context.Context, login string, password string) (int, error) {
	UserID, err := s.Repo.Create(ctx, login, password)
	if err != nil {
		return 0, err
	}
	return UserID, nil
}
