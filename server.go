package main

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		return
	}
	err = client.Ping(ctx, readpref.Primary())

	// root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (b bool, e error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			println("request to /users")
			return next(context)
		}
	}

	// home routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// params - middleware is used
	e.GET("/users/:id", getUser, track)
	// form data
	e.POST("/save", save)
	// file save
	e.POST("/save_photo", savePhoto)
	// query params
	e.GET("/show", show)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func save(c echo.Context) error {
	// using the struct for json parse
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func savePhoto(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")

	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")

}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+", member: "+member)
}
