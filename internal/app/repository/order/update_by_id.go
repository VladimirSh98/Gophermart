package order

import (
	"context"
	"database/sql"
)

func (repo *Repository) UpdateByID(ctx context.Context, orderID string, status string, value sql.NullFloat64) error {
	query := "UPDATE \"order\" SET status = $1, value = $2 WHERE id = $3"
	_, err := repo.Conn.Exec(query, status, value, orderID)
	if err != nil {
		return err
	}
	return nil
}
