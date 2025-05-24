package operation

import (
	"database/sql"
	"time"
)

type Repository struct {
	Conn *sql.DB
}

type Operation struct {
	ID        string
	CreatedAt time.Time
	UserID    int
	Value     float64
}
