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

// V1Cat Route '/v1/cats' to handler function
func V1Cat(g *echo.Group) {
	h := new(handlers.CatHandler)
	g.GET("/:id", h.GetOne)
	g.GET("", h.GetMore)
	g.POST("", h.Add)
	g.DELETE("/:id", h.Delete)
	// g.PATCH("/:id", h.Update)
}

func V2Cat(g *echo.Group) {
	h := new(handlers.CatHandler)

	g.GET("", h.V2GetMore)
}

// V1CatComment Route'/v1/cats/:id/comments'
func V1CatComment(g *echo.Group) {
	h := new(handlers.CatCommentHandler)
	g.GET("/:cid", h.GetOne)
	g.GET("", h.GetMore)
	g.POST("", h.Add)
	g.DELETE("/:cid", h.Delete)
	// g.PATCH("/:id/comments/:cid", h.Update)
}
