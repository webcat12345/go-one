package api

import (
	"github.com/labstack/echo"
	"github.com/webcat12345/go-one/core/server"
	"github.com/webcat12345/go-one/core/services"
	"net/http"
	"strconv"
)

type userHandler struct {
	userService services.UserService
}

func MountUserHandler(group *echo.Group, service services.UserService) {

	handler := userHandler{
		userService: service,
	}

	group.GET("/users", handler.getUsers)
	group.GET("/users/:id", handler.getUserById)
	group.POST("/users", handler.createUser)
}

func (h *userHandler) getUserById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get id from the request")
	}
	user, err := h.userService.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get user from id")
	}
	return ctx.JSON(http.StatusOK, server.JSON{
		Success: true,
		Data:    user,
	})
}

func (h *userHandler) getUsers(ctx echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, server.JSON{
		Success: true,
		Data:    users,
	})
}

func (h *userHandler) createUser(ctx echo.Context) error {
	var payload struct {
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required,length(8|64)"`
	}
	err := ctx.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse payload")
	}
	user, err := h.userService.CreateUser(payload.Email, payload.Password)

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, server.JSON{
		Success: true,
		Data:    user,
	})
}
