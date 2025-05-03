package user

import (
	"database/sql"
)

func (repo *Repository) Create(login string, password string) (sql.Result, error) {
	query := "INSERT INTO \"user\" (login, password, archived) VALUES ($1, $2, false);"
	res, err := repo.Conn.Exec(query, login, password)

	if err != nil {
		return nil, err
	}
	return res, nil
}
