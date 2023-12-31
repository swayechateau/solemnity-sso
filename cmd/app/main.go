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
