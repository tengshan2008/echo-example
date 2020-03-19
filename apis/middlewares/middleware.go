package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Auth Info
const (
	Username = "admin"
	Password = "1234"
)

// Custom Middleware

// SetAdminMiddlewares middleware adds a `logger` to the admin
func SetAdminMiddlewares(g *echo.Group) {
	// logs all server interaction
	loggerConfig := middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}]` + "\n",
	}
	g.Use(middleware.LoggerWithConfig(loggerConfig))

	// Basic Authentication
	bavFunc := func(username, password string, c echo.Context) (bool, error) {
		if username == Username && password == Password {
			return true, nil
		}
		return false, nil
	}
	g.Use(middleware.BasicAuth(bavFunc))
}

// ServerHeader middleware adds a `server` header to the response
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}

// SetMainMiddlewares main
// use custom middleware
func SetMainMiddlewares(e *echo.Echo) {
	e.Use(serverHeader)
}
