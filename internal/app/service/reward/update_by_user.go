package reward

func (s *Service) UpdateByUser(UserID int, balance float64, withdrawn float64) error {
	err := s.Repo.UpdateByUser(UserID, balance, withdrawn)
	if err != nil {
		return err
	}
	return nil
}
