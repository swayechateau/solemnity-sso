package auth

import (
	"context"
	"fmt"
	"net/http"
	"sso/auth/config"
	"sso/auth/microsoft"
	"sso/auth/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var microsoftOauthConfig = &oauth2.Config{
	RedirectURL:  config.GetMicrosoftConfig().RedirectUrl,
	ClientID:     config.GetMicrosoftConfig().ClientId,
	ClientSecret: config.GetMicrosoftConfig().ClientSecret,
	Scopes:       []string{microsoft.Scopes.Profile},
	Endpoint: oauth2.Endpoint{
		AuthURL:  microsoft.AzureADEndpoint("").AuthURL,
		TokenURL: microsoft.AzureADEndpoint("").TokenURL,
	},
}

func MicrosoftLoginHandler(c echo.Context) error {
	url := microsoftOauthConfig.AuthCodeURL(randomStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func MicrosoftCallbackHandler(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")

	if state != randomStateString {
		return fmt.Errorf("invalid oauth state")
	}

	token, err := microsoftOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange failed: %s", err.Error())
	}

	content, err := user.GetOAuthInfo(token, microsoft.Api)
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.String(http.StatusOK, "Content: "+string(content))
}

func MicrosoftHandler(a *echo.Group) {
	a.GET("/microsoft", MicrosoftLoginHandler)
	a.GET("/microsoft/callback", MicrosoftCallbackHandler)
}
