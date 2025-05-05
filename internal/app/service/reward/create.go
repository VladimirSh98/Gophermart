package reward

func (s *Service) Create(UserID int) error {
	_, err := s.Repo.Create(UserID)
	if err != nil {
		return err
	}
	return nil
}
