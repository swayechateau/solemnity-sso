package auth

import (
	"fmt"

	"github.com/swayedev/way"
)

func QRCodeLoginHandler(c *way.Context) {
	referer := "me"
	c.Response.Write([]byte(fmt.Sprintf("referer: %v", referer)))
}
