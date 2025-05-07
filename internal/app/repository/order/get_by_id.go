package order

func (repo *Repository) GetByID(OrderID string) (Order, error) {
	var record Order
	query := "SELECT * FROM \"order\" WHERE id = $1"
	row := repo.Conn.QueryRow(query, OrderID)
	err := row.Scan(&record.ID, &record.UserID, &record.Status, &record.Value, &record.UploadedAt)
	if err != nil {
		return record, err
	}
	return record, nil
}
