package auth

import (
	"fmt"
	"log"
	"net/http"
	"sso/internal/config"
	"sso/pkg/oauth2/github"
	"sso/pkg/oauth2/provider"

	"github.com/swayedev/way"

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

var githubStateString string

func GithubLoginHandler(c *way.Context) {
	githubStateString = randomCode()
	url := githubOauthConfig.AuthCodeURL(githubStateString)
	c.Redirect(url, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(c *way.Context) {
	state := c.Request.FormValue("state")
	code := c.Request.FormValue("code")
	ctx := c.Request.Context()

	if state != githubStateString {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid oauth state"))
	}

	token, err := githubOauthConfig.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("code exchange failed: %s", err.Error()))
	}

	content, err := provider.GetOAuthInfo(github.TokenType, token, github.Api)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect("/", http.StatusTemporaryRedirect)
	}

	c.JSON(http.StatusOK, "Content: "+string(content))
}
