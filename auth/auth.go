package auth

import (
	"sso/database"

	"github.com/labstack/echo/v4"
)

var randomStateString = "state" // make([]byte, 16)
var db *AuthDB

func SetDb(conn *database.Conn) {
	db = NewAuthDBHandler(conn)
}

func AuthHandler(e *echo.Echo) {
	a := e.Group("/auth")
	GoogleHandler(a)
	GithubHandler(a)
	MicrosoftHandler(a)
	// QRCodeHandler(a)

	UserHandler(a)
}
