package reward

import "time"

func (repo *Repository) AccrueReward(userID int, accrual float64) error {
	query := "UPDATE \"reward\" SET balance = balance + $1, updated_at = $2 WHERE user_id = $3"
	_, err := repo.Conn.Exec(query, accrual, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
