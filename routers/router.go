package routers

import (
	"echo-example/apis"

	"github.com/labstack/echo"
)

// New Echo
func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// use debug mode, could config this.
	e.Debug = true

	// set main routes
	apis.MainGroup(e)

	//set admin group routes
	apis.AdminGroup(e.Group("/admin"))

	// set cat group routes
	apis.CatGroup(e.Group("/v1/cats"))

	return e
}
