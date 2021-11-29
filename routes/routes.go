package routes

import (
	"antonio/factory"
	"antonio/config"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	presenter:=factory.Init()

	conn := echo.New()
	jwt := conn.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_SECRET)))
	conn.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	
	jwt.GET("/users", presenter.UserPresentation.GetUsersHandler)
	jwt.GET("/users/:id", presenter.UserPresentation.GetUserByIdHandler)
	conn.POST("/users/register", presenter.UserPresentation.RegisterUserHandler)
	conn.POST("/users/login", presenter.UserPresentation.LoginUserHandler)

	return conn

}