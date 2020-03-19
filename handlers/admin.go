package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// MainAdmin admin handler
func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "You are on the Admin Main Page !!!")
}
