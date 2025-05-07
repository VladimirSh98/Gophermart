package order

func (repo *Repository) UpdateByID(OrderID string, Status string) error {
	query := "UPDATE \"order\" SET status = $1 WHERE id = $2"
	_, err := repo.Conn.Exec(query, Status, OrderID)
	if err != nil {
		return err
	}
	return nil
}
