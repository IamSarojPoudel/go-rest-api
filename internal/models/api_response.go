package models

type APIResponse struct {
	Type       string      `json:"type"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
