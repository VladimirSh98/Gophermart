package operation

import "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operation"

type ServiceInterface interface {
	GetByUser(UserID int) ([]operation.Operation, error)
}

type Service struct {
	Repo operation.Repository
}

func NewService(repo operation.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
