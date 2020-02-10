package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello ...
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
