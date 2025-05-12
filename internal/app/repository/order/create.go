package order

import "database/sql"

func (repo *Repository) Create(orderID string, userID int) (sql.Result, error) {
	query := "INSERT INTO \"order\" (id, user_id, status) VALUES ($1, $2, 'NEW');"
	res, err := repo.Conn.Exec(query, orderID, userID)

	if err != nil {
		return nil, err
	}
	return res, nil
}
