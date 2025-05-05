package reward

import (
	rewardRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
)

func (s *Service) GetByUser(UserID int) (rewardRepo.Reward, error) {
	reward, err := s.Repo.GetByUser(UserID)
	if err != nil {
		return reward, err
	}
	return reward, nil
}
