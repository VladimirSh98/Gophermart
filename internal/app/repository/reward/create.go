package reward

import "database/sql"

func (repo *Repository) Create(UserID int) (sql.Result, error) {
	query := "INSERT INTO \"reward\" (user_id) VALUES ($1);"
	res, err := repo.Conn.Exec(query, UserID)

	if err != nil {
		return nil, err
	}
	return res, nil
}
