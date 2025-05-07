package order

func (s *Service) Create(OrderID string, UserID int) error {
	_, err := s.Repo.Create(OrderID, UserID)
	if err != nil {
		return err
	}
	return nil
}
