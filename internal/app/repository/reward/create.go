package reward

import (
	"context"
	"database/sql"
)

func (repo *Repository) Create(ctx context.Context, userID int) (sql.Result, error) {
	query := "INSERT INTO \"reward\" (user_id) VALUES ($1);"
	res, err := repo.Conn.Exec(query, userID)

	if err != nil {
		return nil, err
	}
	return res, nil
}
