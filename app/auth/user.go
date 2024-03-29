package auth

import (
	"context"
	"fmt"
	"net/http"
	"sso/app/auth/user"
	"sso/database"
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

func AddMe(conn *database.Conn, ctx context.Context) error {
	uuid := models.UUIDStringToBytes("7ccbce2f-3654-4497-8f62-7e11b89e98ce")
	pId := "990079e3-8327-4bf0-9ede-28e27ab22a9b"

	fmt.Printf("Uuid: %v", uuid)
	nU := models.User{
		Id:               uuid,
		Verified:         true,
		DisplayName:      "Swaye Chateau",
		PrimaryEmail:     "swaye@dev.com",
		PrimaryPictureId: pId,
		PrimaryLanguage:  "en",
	}

	err := conn.CreateUser(ctx, &nU)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
