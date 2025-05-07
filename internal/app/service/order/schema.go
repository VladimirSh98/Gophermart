package order

import (
	"database/sql"
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

type ServiceInterface interface {
	GetByUser(UserID int) ([]orderRepo.Order, error)
	GetByID(OrderID string) (orderRepo.Order, error)
	Create(OrderID string, UserID int) error
	UpdateByID(OrderID string, Status string, Value sql.NullFloat64) error
}

type Service struct {
	Repo orderRepo.Repository
}

func NewService(repo orderRepo.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
