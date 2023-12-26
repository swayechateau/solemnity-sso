package auth

import (
	"context"
	"fmt"
	"net/http"
	"sso/auth/config"
	"sso/auth/google"
	"sso/auth/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  config.GetGoogleConfig().RedirectUrl,
	ClientID:     config.GetGoogleConfig().ClientId,
	ClientSecret: config.GetGoogleConfig().ClientSecret,
	Scopes:       []string{google.Scopes.Profile, google.Scopes.Email},
	Endpoint:     google.Endpoint,
}

func GoogleLoginHandler(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(randomStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallbackHandler(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")

	if state != randomStateString {
		return fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange failed: %s", err.Error())
	}

	content, err := user.GetOAuthInfo(token, google.Api)
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.String(http.StatusOK, "Content: "+string(content))
}

func GoogleHandler(a *echo.Group) {
	a.GET("/google", GoogleLoginHandler)
	a.GET("/google/callback", GoogleCallbackHandler)
}
