package reward

func (s *Service) Create(userID int) error {
	_, err := s.Repo.Create(userID)
	if err != nil {
		return err
	}
	return nil
}
