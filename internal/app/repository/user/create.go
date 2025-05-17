package user

import "context"

func (repo *Repository) Create(ctx context.Context, login string, password string) (int, error) {
	query := "INSERT INTO \"user\" (login, hash, archived) VALUES ($1, $2, false) RETURNING id;"
	var ID int
	err := repo.Conn.QueryRow(query, login, password).Scan(&ID)

	if err != nil {
		return 0, err
	}
	return ID, nil
}
