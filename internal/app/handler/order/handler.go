package order

import (
	accrualService "github.com/VladimirSh98/Gophermart.git/internal/app/service/accrual"
	orderService "github.com/VladimirSh98/Gophermart.git/internal/app/service/order"
	rewardService "github.com/VladimirSh98/Gophermart.git/internal/app/service/reward"
)

type Handler struct {
	Order   orderService.ServiceInterface
	Accrual accrualService.ServiceInterface
	Reward  rewardService.ServiceInterface
}

func NewHandler(
	order orderService.ServiceInterface,
	accrual accrualService.ServiceInterface,
	Reward rewardService.ServiceInterface,
) *Handler {
	return &Handler{Order: order, Accrual: accrual, Reward: Reward}
}
