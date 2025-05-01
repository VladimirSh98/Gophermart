package order

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

type ServiceInterface interface{}

type Service struct {
	Repo order.Repository
}

func NewService(repo order.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
