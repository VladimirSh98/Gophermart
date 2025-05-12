package accrual

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *HTTPClient) GetByNumber(ctx context.Context, number string) (*Calculations, error) {
	var data Calculations
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/orders/%s", c.baseURL, number), nil)
	if err != nil {
		return &data, err
	}
	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return &data, err
	}
	defer resp.Body.Close()
	data.StatusCode = resp.StatusCode
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
