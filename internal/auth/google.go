package auth

import (
	"fmt"
	"log"
	"net/http"
	"sso/internal/config"
	"sso/internal/database"
	"sso/internal/database/models"

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
	found, err := findOrUpdateUser(c, "google", context.Id, "", context.Email)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	if found {
		return
	}
	// create new user
	log.Printf("User not found, creating user")
	u := models.NewUser()
	u.SetDisplayName(context.Name)
	u.SetPrimaryEmail(context.Email)
	u.PrimaryLanguage = "en"
	u.Verified = true
	if err := database.CreateUser(c, u); err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	_, err = findOrUpdateUser(c, "google", context.Id, "", context.Email)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
}

func findOrUpdateUser(c *way.Context, providerName string, providerId string, principal string, email string) (bool, error) {
	id, err := database.FindUserIdByProvider(c, "google", c.HashStringToString(providerId))
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return false, err
	}

	if id != nil {
		log.Printf("User found, sending user")
		u, err := database.GetUser(c, [16]byte(id))
		if err != nil {
			log.Printf("ERROR: %v", err.Error())
			c.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
			return false, err
		}
		c.JSON(http.StatusOK, u.ToJson())
		return true, nil
	}

	id, err = database.FindUserIdByEmail(c, c.HashStringToString(email))
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return false, err
	}

	if id != nil {
		log.Printf("User found, sending user")
		u, err := database.GetUser(c, [16]byte(id))
		if err != nil {
			log.Printf("ERROR: %v", err.Error())
			c.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
			return false, err
		}

		if principal == "" {
			principal = email
		}
		if err := database.CreateProvider(c, models.SetProvider("google", principal, providerId, u.Id)); err != nil {
			log.Printf("ERROR: %v", err.Error())
			c.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
			return false, err
		}

		c.JSON(http.StatusOK, u.ToJson())
		return true, nil
	}

	return false, nil
}
