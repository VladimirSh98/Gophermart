package user

import "context"

func (repo *Repository) GetUserByLogin(ctx context.Context, login string, archived bool) (User, error) {
	var record User
	query := "SELECT * FROM \"user\" WHERE login = $1 and archived = $2"
	row := repo.Conn.QueryRow(query, login, archived)
	err := row.Scan(&record.ID, &record.CreatedAt, &record.Login, &record.Password, &record.Archived)
	if err != nil {
		return record, err
	}
	return record, nil
}
