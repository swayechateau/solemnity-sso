package main

import (
	"net/http"
	"sso/app"
	"sso/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	// auth routes
	auth.AuthHandler(e)
	e.GET("/", yourHandler)

	port := app.Port()

	e.Start(":" + port)
}

func yourHandler(c echo.Context) error {
	// Get the Referer header
	referer := c.Request().Header.Get("Referer")

	return c.String(http.StatusOK, referer)
}
