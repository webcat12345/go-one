package route

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/webcat12345/go-one/api"
	"github.com/webcat12345/go-one/core/server"
	"github.com/webcat12345/go-one/core/services"
	"gopkg.in/go-playground/validator.v9"
)

func Init(db *pg.DB) *echo.Echo {
	e := echo.New()

	// root middleware registration
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	// TODO: Custom CORS setting with CORSWithConfig
	e.Use(middleware.CORS())

	e.Validator = &server.CustomValidator{
		OneValidator: validator.New(),
	}

	// create service instance
	userService := services.NewUserService(db)

	// route config
	v1 := e.Group("/api/v1")
	api.MountUserHandler(v1, userService)

	return e
}
