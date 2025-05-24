package reward

import "context"

func (s *Service) AccrueReward(ctx context.Context, userID int, accrual float64) error {
	err := s.Repo.AccrueReward(ctx, userID, accrual)
	if err != nil {
		return err
	}
	return nil
}
