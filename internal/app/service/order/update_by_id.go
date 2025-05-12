package order

import "database/sql"

func (s *Service) UpdateByID(orderID string, status string, value sql.NullFloat64) error {
	err := s.Repo.UpdateByID(orderID, status, value)
	if err != nil {
		return err
	}
	return nil
}
