package server

type JSON struct {
	Status  int               `json:"status"`
	Data    interface{}       `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}
