package main

import (
	"echo-gorm-example/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Bind(e)

	e.Logger.Fatal(e.Start(":1323"))
}
