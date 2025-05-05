package user

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
)

type ServiceInterface interface {
	GetByLogin(login string, archived bool) (user.User, error)
	Create(login string, password string) (int, error)
}

type Service struct {
	Repo user.Repository
}

func NewService(repo user.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
