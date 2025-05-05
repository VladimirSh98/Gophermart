package user

import userRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"

func (s *Service) GetByLogin(login string, archived bool) (userRepo.User, error) {
	user, err := s.Repo.GetUserByLogin(login, archived)
	if err != nil {
		return user, err
	}
	return user, nil
}
