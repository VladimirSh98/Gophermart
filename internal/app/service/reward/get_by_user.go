package reward

import (
	"context"
	rewardRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
)

func (s *Service) GetByUser(ctx context.Context, userID int) (rewardRepo.Reward, error) {
	reward, err := s.Repo.GetByUser(ctx, userID)
	if err != nil {
		return reward, err
	}
	return reward, nil
}
