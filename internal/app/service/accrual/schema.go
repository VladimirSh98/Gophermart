package accrual

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/client/accrual"
)

type Service struct {
	Client accrual.HTTPClient
}

func NewService(client accrual.HTTPClient) ServiceInterface {
	return &Service{Client: client}
}
