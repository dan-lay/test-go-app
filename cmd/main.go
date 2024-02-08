package main

import (
	"fmt"
	"test-go-app/handler"

	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	var db DB

	userHandler := handler.UserHandler{db}

	app.Start(":3000")

	fmt.Println("hello it's working")
}