package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetResources(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Success")
}
