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
	// use debug mode, could config this.
	e.Debug = true

	middlewares.SetLogMiddlewares(e)

	// set main routes
	apis.MainGroup(e)

	//set admin group routes
	apis.AdminGroup(e.Group("/admin"))

	// set v1 group routes
	v1 := e.Group("/v1")

	// set cat group routes
	apis.V1Cat(v1.Group("/cats"))

	// set cat comment group routes
	apis.V1CatComment(v1.Group("/:id/comments"))

	// set v2 group routes
	v2 := e.Group("/v2")

	apis.V2Cat(v2.Group("/cats"))

	return e
}
