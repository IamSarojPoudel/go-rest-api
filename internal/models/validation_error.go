package models

type ValidationErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
