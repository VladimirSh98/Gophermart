package order

import (
	"context"
	"database/sql"
)

func (s *Service) UpdateByID(ctx context.Context, orderID string, status string, value sql.NullFloat64) error {
	err := s.Repo.UpdateByID(ctx, orderID, status, value)
	if err != nil {
		return err
	}
	return nil
}
