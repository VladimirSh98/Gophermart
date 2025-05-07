package order

import "database/sql"

func (s *Service) UpdateByID(OrderID string, Status string, Value sql.NullFloat64) error {
	err := s.Repo.UpdateByID(OrderID, Status, Value)
	if err != nil {
		return err
	}
	return nil
}
