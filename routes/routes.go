package routes

import (
	"github.com/labstack/echo/middleware"
	om "mpbe/middleware"
	"github.com/labstack/echo"
)

func New() *echo.Echo{
	e:=echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	om.LogMiddleware(e)

	e.GET()
	
	return e
}