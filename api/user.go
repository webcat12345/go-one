package api

import (
	"github.com/labstack/echo"
	"github.com/webcat12345/go-one/core/server"
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
	group.POST("/user", handler.createUser)
}

func (h *userHandler) getUsers(ctx echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, server.JSON{
			Status:  http.StatusBadRequest,
			Message: "Failed to get users",
		})
	}
	return ctx.JSON(http.StatusOK, server.JSON{
		Status: http.StatusOK,
		Data:   users,
	})
}

func (h *userHandler) createUser(ctx echo.Context) error {
	var payload struct {
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required,length(8|64)"`
	}
	err := ctx.Bind(&payload)
	if err != nil {
		return nil
	}
	user, err := h.userService.CreateUser(payload.Email, payload.Password)
	return ctx.JSON(http.StatusCreated, server.JSON{
		Status: http.StatusCreated,
		Data:   user,
	})
}
