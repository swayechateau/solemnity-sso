package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"sso/app"
	"sso/github"
	"sso/microsoft"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  app.Url("/callback"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	microsoftOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		RedirectURL:  app.Url("/callback"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     microsoft.AzureADEndpoint("TENANT_ID"),
	}

	githubOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  app.Url("/callback"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     github.Endpoint,
	}

	// Apple login requires additional setup with private keys and configurations.
	// Refer to Apple's documentation for implementation details.
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", handleLogin)
	e.GET("/callback", handleCallback)

	port := app.Port()

	e.Start(":" + port)
}

func handleLogin(c echo.Context) error {
	googleURL := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	microsoftURL := microsoftOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	githubURL := microsoftOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	// Construct Apple login URL based on your implementation.

	data := map[string]string{
		"google_login_url":    googleURL,
		"microsoft_login_url": microsoftURL,
		"github_login_url":    githubURL,
		// Add "apple_login_url" to the data map.
	}

	return c.Render(http.StatusOK, "login.html", data)
}

func handleCallback(c echo.Context) error {
	provider := c.QueryParam("provider")
	code := c.QueryParam("code")
	var token *oauth2.Token
	var err error

	switch provider {
	case "google":
		token, err = googleOauthConfig.Exchange(context.Background(), code)
	case "microsoft":
		token, err = microsoftOauthConfig.Exchange(context.Background(), code)
	case "github":
		token, err = githubOauthConfig.Exchange(context.Background(), code)
		// Add a case for "apple".
	}

	if err != nil {
		fmt.Println("Error exchanging code for token:", err)
		return c.String(http.StatusInternalServerError, "Error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token.AccessToken,
		"token_type":   token.TokenType,
	})
}
