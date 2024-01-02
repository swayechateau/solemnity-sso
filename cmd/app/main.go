package app

import (
	"fmt"
	"sso/internal/auth"
	"sso/internal/config"
	"sso/internal/demo"

	"github.com/swayedev/way"
)

func App() {
	way := way.New()
	if err := way.Db().PgxOpen(); err != nil {
		panic(err)
	}

	auth.AuthHandler(way)
	way.GET("/", yourHandler)
	way.GET("/user", demo.AddUserHandler)

	port := config.GetPort()
	way.Start(":" + port)
}

func yourHandler(c *way.Context) {
	// Get the Referer header
	referer := "me"
	c.Response.Write([]byte(fmt.Sprintf("referer: %v", referer)))
}
