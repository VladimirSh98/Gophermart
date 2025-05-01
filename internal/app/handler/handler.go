package handler

import (
	operationsService "github.com/VladimirSh98/Gophermart.git/internal/app/service/operations"
	orderService "github.com/VladimirSh98/Gophermart.git/internal/app/service/order"
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
	userService "github.com/VladimirSh98/Gophermart.git/internal/app/service/user"
)

type Handler struct {
	operations operationsService.ServiceInterface
	order      orderService.ServiceInterface
	reward     rewardService.ServiceInterface
	user       userService.ServiceInterface
}

func NewHandler(
	operations operationsService.ServiceInterface,
	order orderService.ServiceInterface,
	reward rewardService.ServiceInterface,
	user userService.ServiceInterface,
) *Handler {
	return &Handler{operations: operations, order: order, reward: reward, user: user}
}
