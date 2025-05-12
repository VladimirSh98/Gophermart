package order

import "context"

func (repo *Repository) GetByUser(ctx context.Context, userID int) ([]Order, error) {
	results := make([]Order, 0)
	query := "SELECT * FROM \"order\" WHERE user_id = $1 ORDER BY uploaded_at DESC"
	rows, err := repo.Conn.Query(query, userID)
	if err != nil {
		return results, err
	}
	err = rows.Err()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var record Order
		err = rows.Scan(&record.ID, &record.UserID, &record.Status, &record.Value, &record.UploadedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}
	return results, nil
}
