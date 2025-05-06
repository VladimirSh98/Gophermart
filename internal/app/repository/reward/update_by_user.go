package reward

import "time"

func (repo *Repository) UpdateByUser(UserID int, balance float64, withdrawn float64) error {
	query := "UPDATE \"reward\" SET balance = $1, withdrawn = $2, updated_at = $3 WHERE user_id = $4"
	_, err := repo.Conn.Exec(query, balance, withdrawn, time.Now(), UserID)
	if err != nil {
		return err
	}
	return nil
}
