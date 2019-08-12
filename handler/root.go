package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BindRootHandler(e *echo.Echo) {
	e.GET("/", showIndexPage)
}

func showIndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "API is not implemented yet.")
}
