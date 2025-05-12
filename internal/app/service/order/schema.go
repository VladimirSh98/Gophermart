package order

import (
	"database/sql"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

type ServiceInterface interface {
	GetByUser(userID int) ([]orderRepo.Order, error)
	GetByID(orderID string) (orderRepo.Order, error)
	Create(orderID string, userID int) error
	UpdateByID(orderID string, status string, value sql.NullFloat64) error
}

type Service struct {
	Repo orderRepo.Repository
}

func NewService(repo orderRepo.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
