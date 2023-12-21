package main

import (
	"context"

	"github.com/Maxime-Hrt/got/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/public", "public")

	userHandler := handler.UserHandler{}
	app.Use(withUser)
	app.GET("/user", userHandler.HandlerUserShow)

	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "test@example.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
