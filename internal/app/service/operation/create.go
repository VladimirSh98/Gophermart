package operation

import "context"

func (s *Service) Create(ctx context.Context, orderID string, userID int, Value float64) error {
	_, err := s.Repo.Create(ctx, orderID, userID, Value)
	if err != nil {
		return err
	}
	return nil
}
