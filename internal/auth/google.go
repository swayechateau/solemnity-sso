package auth

import (
	"fmt"
	"log"
	"net/http"
	"sso/internal/config"
	"sso/internal/database"
	"sso/pkg/oauth2/google"
	"sso/pkg/oauth2/provider"

	"github.com/swayedev/way"

	"golang.org/x/oauth2"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  config.GetGoogleConfig().RedirectUrl,
	ClientID:     config.GetGoogleConfig().ClientId,
	ClientSecret: config.GetGoogleConfig().ClientSecret,
	Scopes:       []string{google.Scopes.Profile, google.Scopes.Email},
	Endpoint:     google.Endpoint,
}

var googleStateString string

func GoogleLoginHandler(c *way.Context) {
	googleStateString = randomCode()
	url := googleOauthConfig.AuthCodeURL(googleStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallbackHandler(c *way.Context) {
	state := c.Request.URL.Query().Get("state")
	code := c.Request.URL.Query().Get("code")
	ctx := c.Request.Context()

	if state != googleStateString {
		// Create a JSON response with the error message
		errMsg := "invalid oauth state"
		c.JSON(http.StatusBadRequest, ErrorMessage{errMsg})
		return
	}

	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		// Format the error message
		errMsg := fmt.Sprintf("code exchange failed: %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorMessage{errMsg})
		return
	}

	content, err := provider.GetOAuthInfo(google.TokenType, token, google.Api)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("Content: %v", content)
	context := google.JsonToContext(content)

	id, err := database.FindUserIdByProvider(c, "google", context.Id)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
	}

	if id != nil {
		log.Printf("User found, sending user")
		u, err := database.GetUser(c, [16]byte(id))
		if err != nil {
			log.Printf("ERROR: %v", err.Error())
		}
		c.JSON(http.StatusOK, u.ToJson())
	}

	id, err = database.FindUserIdByEmail(c, context.Email)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
	}

	if id != nil {
		log.Printf("User found, sending user")
		u, err := database.GetUser(c, [16]byte(id))
		if err != nil {
			log.Printf("ERROR: %v", err.Error())
		}
		c.JSON(http.StatusOK, u.ToJson())
	}

	log.Printf("User not found, sending context")
	c.JSON(http.StatusOK, context)
}
