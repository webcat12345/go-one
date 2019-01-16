package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/webcat12345/go-one/api"
	"github.com/webcat12345/go-one/db"
)

func Init() *echo.Echo {
	e := echo.New()

	// root middleware registration
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	// TODO: Custom CORS setting with CORSWithConfig
	e.Use(middleware.CORS())

	// db connection check
	db := db.Init()

	collection := db.Database("go-one").Collection("assets")
	println(collection)

	// route config
	v1 := e.Group("/api/v1")
	v1.GET("/all", api.GetResources)

	return e
}
