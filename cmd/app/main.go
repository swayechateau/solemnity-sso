package app

import (
	"fmt"
	"sso/internal/auth"
	"sso/internal/config"

	"github.com/swayedev/way"
)

func App() {
	way := way.New()
	// w.Use(middleware.Logger())
	// conn, err := way.Open()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	auth.AuthHandler(way)
	way.GET("/", yourHandler)
	// way.GET("/auth/google", auth.GoogleLoginHandler)
	// way.GET("/auth/google/callback", auth.GoogleCallbackHandler)

	// way.GET("/auth/github", auth.GithubLoginHandler)
	// way.GET("/auth/github/callback", auth.GithubCallbackHandler)

	way.GET("/auth/microsoft", auth.MicrosoftLoginHandler)
	way.GET("/auth/microsoft/callback", auth.MicrosoftCallbackHandler)

	// way.GET("/oauth2/qr", auth.QRCodeLoginHandler)
	// w.GET("/user/:id", dbHandler(c))

	port := config.GetPort()
	way.Start(":" + port)
}

func yourHandler(c *way.Context) {
	// Get the Referer header
	referer := "me"
	c.Response.Write([]byte(fmt.Sprintf("referer: %v", referer)))
}

// func dbHandler(c way.Context) error {
// 	// Get the Referer header
// 	id := c.Param("id")
// 	ctx := c.Request().Context()
// 	fmt.Print(id)
// 	// Find user by id
// 	u, err := conn.FindUserById(ctx, models.UUIDStringToBytes(id))
// 	if err != nil {
// 		fmt.Print(err)
// 		return err
// 	}

// 	if u == nil {
// 		err = auth.AddMe(conn, ctx)
// 		if err != nil {
// 			fmt.Print(err)
// 			return err
// 		}

// 		return c.String(http.StatusOK, "User not found, one was created")
// 	}

// 	return c.String(http.StatusOK, fmt.Sprintf("%v", u))
// }
