package presentation

import (
	"net/http"
	"antonio/features/authentication"
	"antonio/features/authentication/presentation/request"
	"antonio/features/authentication/presentation/response"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthBusiness auth.Business
}

func NewAuthHandler(authBusiness auth.Business) *AuthHandler {
	return &AuthHandler{AuthBusiness: authBusiness}
}

func (ah *AuthHandler) LoginHandler(e echo.Context) error {
	user := request.UserRequest{}
	e.Bind(&user)
	accessToken, err := ah.AuthBusiness.Login(user.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "login failed",
			"err":     err.Error(),
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "login success",
		"data": response.AuthResponse{
			AccessToken:  accessToken.Token,
		},
	})
}

