package reward

import "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"

type ServiceInterface interface{}

type Service struct {
	Repo reward.Repository
}

func NewService(repo reward.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
