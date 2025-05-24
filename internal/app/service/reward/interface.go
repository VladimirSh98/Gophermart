package reward

import (
	"context"
	"github.com/VladimirSh98/Gophermart.git/internal/app/repository/reward"
)

type ServiceInterface interface {
	GetByUser(ctx context.Context, userID int) (reward.Reward, error)
	Create(ctx context.Context, userID int) error
	UpdateByUser(ctx context.Context, userID int, balance float64, withdrawn float64) error
	AccrueReward(ctx context.Context, userID int, accrual float64) error
}
