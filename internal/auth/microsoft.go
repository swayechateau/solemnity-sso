package auth

import (
	"fmt"
	"log"
	"net/http"
	"sso/internal/config"
	"sso/pkg/oauth2/microsoft"
	"sso/pkg/oauth2/provider"

	"github.com/swayedev/way"

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

var microsoftStateString string

func MicrosoftLoginHandler(c *way.Context) {
	microsoftStateString = randomCode()
	url := microsoftOauthConfig.AuthCodeURL(microsoftStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func MicrosoftCallbackHandler(c *way.Context) {
	state := c.Request.FormValue("state")
	code := c.Request.FormValue("code")
	ctx := c.Request.Context() // TODO: use context.Background() instead

	if state != microsoftStateString {
		errMsg := "invalid oauth state"
		c.JSON(http.StatusBadRequest, ErrorMessage{errMsg})
		return
	}

	token, err := microsoftOauthConfig.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("code exchange failed: %s", err.Error()))
		return
	}

	content, err := provider.GetOAuthInfo(microsoft.TokenType, token, microsoft.Api)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("Content: %v", content)
	c.JSON(http.StatusOK, microsoft.JsonToContext(content))
}
