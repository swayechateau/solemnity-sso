package main

import (
	"fmt"
	"log"
	"net/http"
	"sso/app/auth"
	"sso/app/config"
	"sso/database"
	"sso/database/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	conn, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	auth.SetDb(conn)
	// auth routes
	auth.AuthHandler(e)

	e.GET("/", yourHandler)
	e.GET("/user/:id", func(c echo.Context) error {
		return dbHandler(conn, c)
	})

	port := config.GetPort()

	e.Start(":" + port)
}

func yourHandler(c echo.Context) error {
	// Get the Referer header
	referer := c.Request().Header.Get("Referer")
	fmt.Print(referer)
	return c.JSON(http.StatusOK, fmt.Sprintf("referer: %v", referer))
}

func dbHandler(conn *database.Conn, c echo.Context) error {
	// Get the Referer header
	id := c.Param("id")
	ctx := c.Request().Context()
	fmt.Print(id)
	// Find user by id
	u, err := conn.FindUserById(ctx, models.UUIDStringToBytes(id))
	if err != nil {
		fmt.Print(err)
		return err
	}

	if u == nil {
		err = auth.AddMe(conn, ctx)
		if err != nil {
			fmt.Print(err)
			return err
		}

		return c.String(http.StatusOK, "User not found, one was created")
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", u))
}
