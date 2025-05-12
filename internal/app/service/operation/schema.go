package operation

import "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"

type ServiceInterface interface {
	GetByUser(userID int) ([]operation.Operation, error)
	Create(orderID string, userID int, value float64) error
}

type Service struct {
	Repo operation.Repository
}

func NewService(repo operation.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
