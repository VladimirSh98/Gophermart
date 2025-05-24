package config

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/models"
	"github.com/caarlos0/env/v6"
)

func parseEnv(conf *models.Config) error {
	err := env.Parse(conf)
	if err != nil {
		return err
	}
	return nil
}
