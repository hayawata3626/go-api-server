package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Person ...
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Hello ...
func Hello(c echo.Context) error {
	// person型のリスト
	persons := []Person{
		{ID: 1, Name: "goro"},
		{ID: 2, Name: "ziro"},
	}
	return c.JSON(http.StatusOK, persons)
}
