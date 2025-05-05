package order

import (
	"database/sql"
	"time"
)

type Repository struct {
	Conn *sql.DB
}

type Order struct {
	ID         string
	UploadedAt time.Time
	UserID     int
	Status     string
	Value      sql.NullFloat64
}
