package handler

import (
	operationService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operation"
	orderService "github.com/VladimirSh98/Gophermart.git/internal/app/service/order"
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
	userService "github.com/VladimirSh98/Gophermart.git/internal/app/service/user"
)

type Handler struct {
	Operation operationService.ServiceInterface
	Order     orderService.ServiceInterface
	Reward    rewardService.ServiceInterface
	User      userService.ServiceInterface
}

func NewHandler(
	operation operationService.ServiceInterface,
	order orderService.ServiceInterface,
	reward rewardService.ServiceInterface,
	user userService.ServiceInterface,
) *Handler {
	return &Handler{Operation: operation, Order: order, Reward: reward, User: user}
}
