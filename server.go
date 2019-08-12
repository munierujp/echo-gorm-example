package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "API is not implemented yet.")
}
