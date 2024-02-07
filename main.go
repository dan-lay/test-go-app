package main

import (
	"context"
	"os"
	"test-go-app/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Main new
	component := templates.Hello("Dan")
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":3000"))
}