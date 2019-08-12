package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ShowIndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "API is not implemented yet.")
}
