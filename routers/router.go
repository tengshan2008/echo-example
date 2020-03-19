package routers

import (
	"echo-example/apis"
	"echo-example/apis/middlewares"

	"github.com/labstack/echo"
)

// New Echo
func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")

	// set middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)

	// set main routes
	apis.MainGroup(e)

	// set group routes
	apis.AdminGroup(adminGroup)

	return e
}
