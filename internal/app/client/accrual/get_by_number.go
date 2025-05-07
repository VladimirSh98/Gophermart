package accrual

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *HTTPClient) GetByNumber(number string) (*Calculations, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/orders/%s", c.baseURL, number), nil)
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data Calculations
	data.StatusCode = resp.StatusCode
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
