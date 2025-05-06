package operation

import (
	operationService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operation"
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
)

type Handler struct {
	Operation operationService.ServiceInterface
	Reward    rewardService.ServiceInterface
}

func NewHandler(operation operationService.ServiceInterface, reward rewardService.ServiceInterface) *Handler {
	return &Handler{Operation: operation, Reward: reward}
}
