package reward

import "context"

func (repo *Repository) GetByUser(ctx context.Context, userID int) (Reward, error) {
	var record Reward
	query := "SELECT * FROM \"reward\" WHERE user_id = $1"
	row := repo.Conn.QueryRow(query, userID)
	err := row.Scan(&record.ID, &record.UserID, &record.Balance, &record.Withdrawn, &record.CreatedAt, &record.UpdatedAt)
	if err != nil {
		return record, err
	}
	return record, nil
}
