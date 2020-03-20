package apis

import (
	"echo-example/handlers"

	"github.com/labstack/echo"
)

// MainGroup Route `/` to handler function
func MainGroup(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheck)
}

// AdminGroup Route '/admin` to handler function
func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}

// CatGroup Route '/cats' to handler function
func CatGroup(g *echo.Group) {
	h := new(handlers.CatHandler)
	g.GET("/:id", h.Get)
	g.POST("", h.Add)
	// g.GET("/:id/comments/:cid", h.GetComments)
	// g.POST("/:id/comments", h.AddComments)
}
