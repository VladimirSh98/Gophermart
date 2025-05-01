package user

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/repository/user"
)

type ServiceInterface interface{}

type Service struct {
	Repo user.Repository
}

func NewService(repo user.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
