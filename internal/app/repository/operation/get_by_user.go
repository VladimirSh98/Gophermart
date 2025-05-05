package operation

func (repo *Repository) GetByUser(UserID int) ([]Operation, error) {
	results := make([]Operation, 0)
	query := "SELECT * FROM \"operation\" WHERE user_id = $1 ORDER BY created_at DESC"
	rows, err := repo.Conn.Query(query, UserID)
	if err != nil {
		return results, err
	}
	err = rows.Err()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var record Operation
		err = rows.Scan(&record.ID, &record.UserID, &record.Value, &record.CreatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}
	return results, nil
}
