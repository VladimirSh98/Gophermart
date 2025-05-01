package operations

import "github.com/VladimirSh98/Gophermart.git/internal/app/repository/operations"

type ServiceInterface interface{}

type Service struct {
	Repo operations.Repository
}

func NewService(repo operations.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
