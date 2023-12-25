package auth

import (
	"github.com/labstack/echo/v4"
)

func QRCodeLoginHandler(c *echo.Context) error {
	return nil
}

func QrHandler(e *echo.Echo) {
	e.GET("/auth/qr", GithubLoginHandler)
}
