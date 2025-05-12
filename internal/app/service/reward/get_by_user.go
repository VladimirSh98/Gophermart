package reward

import (
	rewardRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
)

func (s *Service) GetByUser(userID int) (rewardRepo.Reward, error) {
	reward, err := s.Repo.GetByUser(userID)
	if err != nil {
		return reward, err
	}
	return reward, nil
}
