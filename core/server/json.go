package server

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type (
	JSON struct {
		Success bool              `json:"success"`
		Data    interface{}       `json:"data,omitempty"`
		Message string            `json:"message,omitempty"`
		Errors  map[string]string `json:"errors,omitempty"`
	}

	CustomValidator struct {
		OneValidator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	// TODO: return detailed validation error messages
	return cv.OneValidator.Struct(i)
}

func GetPayloadFromRequest(ctx echo.Context, payload interface{}) error {
	if err := ctx.Bind(payload); err != nil {
		return echo.ErrBadRequest
	}

	if err := ctx.Validate(payload); err != nil {
		return echo.ErrBadRequest
	}
	return nil
}
