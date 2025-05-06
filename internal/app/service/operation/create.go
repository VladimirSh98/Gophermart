package operation

func (s *Service) Create(orderID string, UserID int, Value float64) error {
	_, err := s.Repo.Create(orderID, UserID, Value)
	if err != nil {
		return err
	}
	return nil
}
