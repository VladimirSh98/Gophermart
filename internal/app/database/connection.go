package database

import (
	"database/sql"
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConnectionStruct struct {
	Conn *sql.DB
}

func (db *DBConnectionStruct) OpenConnection(conf *config.Config) error {
	var err error
	db.Conn, err = sql.Open("pgx", conf.DatabaseURI)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBConnectionStruct) CloseConnection() {
	db.Conn.Close()
}
