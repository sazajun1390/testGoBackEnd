package main

import (
	"chatSystem/Infra"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/createUser", Infra.CreateUser)

	e.GET("/signIn", Infra.SignIn)
	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
