package order

import "context"

func (s *Service) Create(ctx context.Context, orderID string, userID int) error {
	_, err := s.Repo.Create(ctx, orderID, userID)
	if err != nil {
		return err
	}
	return nil
}
