package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/webcat12345/go-one/api"
	"github.com/webcat12345/go-one/core/repository"
	"github.com/webcat12345/go-one/core/services"
)

func Init() *echo.Echo {
	e := echo.New()

	// root middleware registration
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	// TODO: Custom CORS setting with CORSWithConfig
	e.Use(middleware.CORS())

	// create db connection
	db := repository.GetDatabase()

	// create service instance
	userService := services.NewUserService(db)

	// route config
	v1 := e.Group("/api/v1")
	api.MountUserHandler(v1, userService)

	return e
}
