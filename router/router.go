package router

import (
	"echo-gorm-example/handler"

	"github.com/labstack/echo/v4"
)

func Bind(e *echo.Echo) {
	handler.BindRootHandler(e)
}
