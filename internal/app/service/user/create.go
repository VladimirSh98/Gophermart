package user

func (s *Service) Create(login string, password string) error {
	_, err := s.Repo.Create(login, password)
	if err != nil {
		return err
	}
	return nil
}
