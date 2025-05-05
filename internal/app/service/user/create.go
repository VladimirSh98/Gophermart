package user

func (s *Service) Create(login string, password string) (int, error) {
	UserID, err := s.Repo.Create(login, password)
	if err != nil {
		return 0, err
	}
	return UserID, nil
}
