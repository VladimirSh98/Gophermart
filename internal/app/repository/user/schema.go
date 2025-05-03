package user

import (
	"database/sql"
	"time"
)

type Repository struct {
	Conn *sql.DB
}

type User struct {
	ID        int
	CreatedAt time.Time
	Login     string
	Password  string
	Archived  bool
}
