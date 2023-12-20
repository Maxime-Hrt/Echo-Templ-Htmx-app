package main

import (
	"github.com/Maxime-Hrt/got/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandlerUserShow)

	app.Start(":3000")
}
