package reward

import "context"

func (s *Service) Create(ctx context.Context, userID int) error {
	_, err := s.Repo.Create(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
