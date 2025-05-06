package order

import (
	accrualService "github.com/VladimirSh98/Gophermart.git/internal/app/service/accrual"
	orderService "github.com/VladimirSh98/Gophermart.git/internal/app/service/order"
)

type Handler struct {
	Order   orderService.ServiceInterface
	Accrual accrualService.ServiceInterface
}

func NewHandler(
	order orderService.ServiceInterface,
	accrual accrualService.ServiceInterface,
) *Handler {
	return &Handler{Order: order, Accrual: accrual}
}
