package auth

import (
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
	userService "github.com/VladimirSh98/Gophermart.git/internal/app/service/user"
)

type Handler struct {
	User   userService.ServiceInterface
	Reward rewardService.ServiceInterface
}

func NewHandler(user userService.ServiceInterface, reward rewardService.ServiceInterface) *Handler {
	return &Handler{User: user, Reward: reward}
}
