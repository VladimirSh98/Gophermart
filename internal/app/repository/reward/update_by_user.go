package reward

import (
	"context"
	"time"
)

func (repo *Repository) UpdateByUser(ctx context.Context, userID int, balance float64, withdrawn float64) error {
	query := "UPDATE \"reward\" SET balance = $1, withdrawn = $2, updated_at = $3 WHERE user_id = $4"
	_, err := repo.Conn.Exec(query, balance, withdrawn, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
