package reward

import "database/sql"

type Repository struct {
	Conn *sql.DB
}
