package auth

import (
	"net/http"
	"sso/auth/user"
	"sso/database/models"

	"github.com/labstack/echo/v4"
)

func GetAuthUserHandler(c echo.Context) error {
	uuid := "7ccbce2f-3654-4497-8f62-7e11b89e98ce"
	id := models.UUIDStringToBytes(uuid)
	ctx := c.Request().Context()
	u, err := user.GetAuthUser(db.Auth, ctx, id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	// Return user details
	return c.JSON(http.StatusOK, u)
}

func UserHandler(a *echo.Group) {
	a.GET("/user", GetAuthUserHandler)
}
