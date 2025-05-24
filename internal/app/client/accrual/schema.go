package accrual

import "net/http"

type HTTPClient struct {
	baseURL    string
	httpClient *http.Client
}

type Calculations struct {
	Order      string  `json:"order"`
	Status     string  `json:"status"`
	Accrual    float64 `json:"accrual"`
	StatusCode int
}
