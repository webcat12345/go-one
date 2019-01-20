package tokenizer

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type JwtAuthClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func JwtCustomConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		Claims:     &JwtAuthClaims{},
		SigningKey: []byte("secret"),
		ErrorHandler: func(e error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		},
	}

	return config
}

func UserIdFromToken(ctx echo.Context) int {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtAuthClaims)
	return claims.Id
}
