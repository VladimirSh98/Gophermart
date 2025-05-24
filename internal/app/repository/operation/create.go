package operation

import (
	"context"
	"database/sql"
)

func (repo *Repository) Create(ctx context.Context, orderID string, userID int, value float64) (sql.Result, error) {
	query := "INSERT INTO \"operation\" (id, user_id, value) VALUES ($1, $2, $3);"
	res, err := repo.Conn.Exec(query, orderID, userID, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
