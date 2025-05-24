package order

import (
	orderRepo "github.com/VladimirSh98/Gophermart.git/internal/app/repository/order"
)

type Service struct {
	Repo orderRepo.Repository
}

func NewService(repo orderRepo.Repository) ServiceInterface {
	return &Service{Repo: repo}
}
