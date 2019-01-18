package api

import (
	"github.com/labstack/echo"
	"github.com/webcat12345/go-one/core/services"
	"net/http"
)

type userHandler struct {
	userService services.UserService
}

func MountUserHandler(group *echo.Group, service services.UserService) {

	handler := userHandler{
		userService: service,
	}

	group.GET("/users", handler.getUsers)
}

func (h *userHandler) getUsers(ctx echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	} else {
		return ctx.JSON(http.StatusOK, &users)
	}
}
