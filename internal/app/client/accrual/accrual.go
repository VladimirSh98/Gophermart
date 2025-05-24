package accrual

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/config"
	"net/http"
	"time"
)

func NewHTTPClient() HTTPClient {
	return HTTPClient{
		baseURL:    config.Conf.AccrualSystemAddress,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}
