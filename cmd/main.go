package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lgomez-dev/GoProject-Template/handler"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}
