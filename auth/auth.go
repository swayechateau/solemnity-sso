package auth

import "github.com/labstack/echo/v4"

var randomStateString = "state" // make([]byte, 16)

func AuthHandler(e *echo.Echo) {
	GoogleHandler(e)
	GithubHandler(e)
	MicrosoftHandler(e)
	// QRCodeHandler(e)
}
