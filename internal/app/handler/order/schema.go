package order

import "time"

var RegisteredStatus = "REGISTERED"
var InvalidStatus = "INVALID"
var ProcessingStatus = "PROCESSING"
var ProcessedStatus = "PROCESSED"

type GetByUserResponse struct {
	Number     string    `json:"number"`
	UploadedAt time.Time `json:"uploaded_at"`
	Accrual    float64   `json:"accrual,omitempty"`
	Status     string    `json:"status"`
}

type ProcessedResult struct {
	OrderID string
	Status  string
	Accrual float64
}
