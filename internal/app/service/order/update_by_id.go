package order

func (s *Service) UpdateByID(OrderID string, Status string) error {
	err := s.Repo.UpdateByID(OrderID, Status)
	if err != nil {
		return err
	}
	return nil
}
