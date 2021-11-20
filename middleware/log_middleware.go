package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time ${time_rfc3339}\nmethod ${method}\nuri ${uri}\nstatus ${status}\nlatency ${latency_human}\n",
	}))

}