package database

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/models"
	"github.com/pressly/goose/v3"
)

func (db *DBConnectionStruct) UpgradeMigrations(conf *models.Config) error {
	err := goose.Up(db.Conn, conf.MigrationsDir)
	if err != nil {
		return err
	}
	return nil
}
