package reward

func (s *Service) AccrueReward(UserID int, accrual float64) error {
	err := s.Repo.AccrueReward(UserID, accrual)
	if err != nil {
		return err
	}
	return nil
}
