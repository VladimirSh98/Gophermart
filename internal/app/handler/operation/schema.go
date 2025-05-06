package operation

import "time"

type GetByUserResponse struct {
	Order       string    `json:"order"`
	Sum         float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

type CreateRequest struct {
	Order string  `json:"order" validate:"required"`
	Sum   float64 `json:"sum" validate:"gte=0"`
}
