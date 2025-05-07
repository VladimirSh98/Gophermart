package order

import "database/sql"

func (repo *Repository) UpdateByID(OrderID string, Status string, Value sql.NullFloat64) error {
	query := "UPDATE \"order\" SET status = $1, value = $2 WHERE id = $3"
	_, err := repo.Conn.Exec(query, Status, Value, OrderID)
	if err != nil {
		return err
	}
	return nil
}
