package operation

func (s *Service) Create(orderID string, userID int, Value float64) error {
	_, err := s.Repo.Create(orderID, userID, Value)
	if err != nil {
		return err
	}
	return nil
}
