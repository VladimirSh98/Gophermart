package reward

import "context"

func (s *Service) UpdateByUser(ctx context.Context, userID int, balance float64, withdrawn float64) error {
	err := s.Repo.UpdateByUser(ctx, userID, balance, withdrawn)
	if err != nil {
		return err
	}
	return nil
}
