package auth

import "time"

type GetByUserResponse struct {
	Number     string    `json:"number"`
	UploadedAt time.Time `json:"uploaded_at"`
	Accrual    float64   `json:"accrual,omitempty"`
	Status     string    `json:"status"`
}
