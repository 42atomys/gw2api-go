package gw2api

import "fmt"

type APIError struct {
	Err  string `json:"error"`
	Text string `json:"text"`
}

func (e APIError) Error() string {
	e.Err += e.Text

	return fmt.Sprintf("API Error: %s", e.Err)
}
