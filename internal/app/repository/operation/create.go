package operation

import "database/sql"

func (repo *Repository) Create(orderID string, UserID int, Value float64) (sql.Result, error) {
	query := "INSERT INTO \"operation\" (id, user_id, value) VALUES ($1, $2, $3);"
	res, err := repo.Conn.Exec(query, orderID, UserID, Value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
