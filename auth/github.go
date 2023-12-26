package auth

import (
	"context"
	"fmt"
	"net/http"

	"sso/auth/config"
	"sso/auth/github"
	"sso/auth/user"

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

	content, err := user.GetOAuthInfo(token, github.Api)
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.String(http.StatusOK, "Content: "+string(content))
}

func GithubHandler(e *echo.Echo) {
	e.GET("/auth/github", GithubLoginHandler)
	e.GET("/auth/github/callback", GithubCallbackHandler)
}