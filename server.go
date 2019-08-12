package main

import (
	"echo-gorm-example/handler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	USER_NAME = "user"
	PASSWORD  = "pass"
	DATABASE  = "sample"
	HOST      = "localhost"
	PORT      = "3306"
	CHARSET   = "utf8mb4"
	COLLATION = "utf8mb4_bin"
)

func main() {
	e := echo.New()

	// Set middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create database connection
	db, err := gorm.Open("mysql", USER_NAME+":"+PASSWORD+"@tcp("+HOST+":"+PORT+")/"+DATABASE+"?charset="+CHARSET+"&collation="+COLLATION+"&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Bind routes
	e.GET("/", h.ShowIndexPage)
	e.POST("/users", h.AddUser)
	e.GET("/users/:id", h.GetUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
