package tokenizer

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func NewAccessToken(userId int) (string, error) {
	claims := &JwtAuthClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
