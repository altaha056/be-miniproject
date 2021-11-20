package middleware

import (
	"mpbe/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"]=userId
	claims["name"]=name
	claims["exp"]=time.Now().Add(time.Hour*1).Unix()

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString([]byte(config.JWTSecret))
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}