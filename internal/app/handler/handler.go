package handler

import (
	operationService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operation"
	orderService "github.com/VladimirSh98/Gophermart.git/internal/app/service/order"
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
	userService "github.com/VladimirSh98/Gophermart.git/internal/app/service/user"
)

type Handler struct {
	operation operationService.ServiceInterface
	order     orderService.ServiceInterface
	reward    rewardService.ServiceInterface
	user      userService.ServiceInterface
}

func NewHandler(
	operation operationService.ServiceInterface,
	order orderService.ServiceInterface,
	reward rewardService.ServiceInterface,
	user userService.ServiceInterface,
) *Handler {
	return &Handler{operation: operation, order: order, reward: reward, user: user}
}
