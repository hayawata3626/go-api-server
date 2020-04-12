package handler

import (
	"log"
	"net/http"

	"go-api-server/repository"

	"github.com/labstack/echo/v4"
)

// ArticleIndex ...
func ArticleIndex(c echo.Context) error {
	// artcle list
	articles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Articles": articles,
	}
	return c.JSON(http.StatusOK, data)
}
