package reward

import (
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
)

type Handler struct {
	Reward rewardService.ServiceInterface
}

func NewHandler(reward rewardService.ServiceInterface) *Handler {
	return &Handler{Reward: reward}
}
