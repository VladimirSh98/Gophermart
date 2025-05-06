package reward

import "github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"

type ServiceInterface interface {
	GetByUser(UserID int) (reward.Reward, error)
	Create(UserID int) error
	UpdateByUser(UserID int, balance float64, withdrawn float64) error
}

type Service struct {
	Repo reward.Repository
}

func NewService(repo reward.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
