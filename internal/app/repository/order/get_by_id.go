package order

import "context"

func (repo *Repository) GetByID(ctx context.Context, orderID string) (Order, error) {
	var record Order
	query := "SELECT * FROM \"order\" WHERE id = $1"
	row := repo.Conn.QueryRow(query, orderID)
	err := row.Scan(&record.ID, &record.UserID, &record.Status, &record.Value, &record.UploadedAt)
	if err != nil {
		return record, err
	}
	return record, nil
}
