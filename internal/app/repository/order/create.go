package order

import "database/sql"

func (repo *Repository) Create(OrderID string, UserID int) (sql.Result, error) {
	query := "INSERT INTO \"order\" (id, user_id, status) VALUES ($1, $2, 'NEW');"
	res, err := repo.Conn.Exec(query, OrderID, UserID)

	if err != nil {
		return nil, err
	}
	return res, nil
}
