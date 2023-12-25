package main

import (
	"fmt"
	"net/http"
	"sso/app"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", yourHandler)

	port := app.Port()

	e.Start(":" + port)
}

func yourHandler(c echo.Context) error {
	// Get the Referer header
	referer := c.Request().Header.Get("Referer")

	fmt.Println(referer)

	// Get client IP address
	clientIP := c.RealIP()

	fmt.Println(clientIP)

	// Now you can use referer and clientIP for your logic
	// ...

	return c.String(http.StatusOK, "Processed request")
}
