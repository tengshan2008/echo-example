package apis

import (
	"echo-example/handlers"

	"github.com/labstack/echo"
)

// MainGroup Route `/` to handler function
func MainGroup(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
}

// AdminGroup Route '/admin` to handler function
func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
