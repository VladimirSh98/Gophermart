package reward

func (s *Service) UpdateByUser(userID int, balance float64, withdrawn float64) error {
	err := s.Repo.UpdateByUser(userID, balance, withdrawn)
	if err != nil {
		return err
	}
	return nil
}
