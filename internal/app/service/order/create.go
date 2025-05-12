package order

func (s *Service) Create(orderID string, userID int) error {
	_, err := s.Repo.Create(orderID, userID)
	if err != nil {
		return err
	}
	return nil
}
