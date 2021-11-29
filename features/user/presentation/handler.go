package presentation

import (
	"antonio/features/user"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(uService user.Service) *UserHandler{
	return &UserHandler{uService}
}

func (uhn *UserHandler)RegisterUserHandler(e echo.Context)error{
}