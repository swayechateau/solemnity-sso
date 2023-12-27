package auth

import (
	"context"
	"fmt"
	"net/http"

	"sso/app/config"
	"sso/oauth2/github"
	"sso/oauth2/provider"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var githubOauthConfig = &oauth2.Config{
	RedirectURL:  config.GetGithubConfig().RedirectUrl,
	ClientID:     config.GetGithubConfig().ClientId,
	ClientSecret: config.GetGithubConfig().ClientSecret,
	Scopes:       []string{github.Scopes.Profile},
	Endpoint: oauth2.Endpoint{
		AuthURL:  github.Endpoint.AuthURL,
		TokenURL: github.Endpoint.TokenURL,
	},
}

func GithubLoginHandler(c echo.Context) error {
	url := githubOauthConfig.AuthCodeURL(randomStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GithubCallbackHandler(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")

	if state != randomStateString {
		return fmt.Errorf("invalid oauth state")
	}

	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange failed: %s", err.Error())
	}

	content, err := provider.GetOAuthInfo(token, github.Api)
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.String(http.StatusOK, "Content: "+string(content))
}

func GithubHandler(a *echo.Group) {
	a.GET("/github", GithubLoginHandler)
	a.GET("/github/callback", GithubCallbackHandler)
}
