package reward

func (s *Service) AccrueReward(userID int, accrual float64) error {
	err := s.Repo.AccrueReward(userID, accrual)
	if err != nil {
		return err
	}
	return nil
}
