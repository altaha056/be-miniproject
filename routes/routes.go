package routes

import (
	"mpbe/factory"
	om "mpbe/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo{
	presenter:=factory.Init()

	e:=echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	om.LogMiddleware(e)

	e.GET("/posts", presenter.PostPresentation.GetAllPost,middleware.BasicAuth(om.BasicAuth))
	
	return e
}