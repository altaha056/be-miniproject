package middleware

import (
	"time"

	"antonio/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateToken(id uint, role string) (string, error) {
	claims := jwt.MapClaims{}

	claims["id"]=id
	claims["limit"]=time.Now().Add(time.Hour * 1).Unix()

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func ExtractClaim(e echo.Context)(claims map[string]interface{}){
	user:=e.Get("user").(*jwt.Token)

	if user.Valid{
		claims=user.Claims.(jwt.MapClaims)
	}
	return
}