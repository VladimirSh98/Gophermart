package reward

import (
	"database/sql"
	"time"
)

type Repository struct {
	Conn *sql.DB
}

type Reward struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	Balance   float64
	Withdrawn float64
}
