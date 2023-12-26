package auth

import (
	"github.com/labstack/echo/v4"
)

func QRCodeLoginHandler(c *echo.Context) error {
	return nil
}

func QrHandler(a *echo.Group) {
	a.GET("/qr", GithubLoginHandler)
}
