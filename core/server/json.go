package server

type JSON struct {
	Success bool              `json:"success"`
	Data    interface{}       `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}
