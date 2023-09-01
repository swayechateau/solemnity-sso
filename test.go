package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		ClientID:     "",
		ClientSecret: "",

		RedirectURL: "http://localhost:8080/callback",
		Scopes:      []string{"openid", "email", "profile"},
		Endpoint:    google.Endpoint,
	}
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", handleGoogleLogin)
	e.GET("/callback", handleGoogleCallback)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Start(":" + port)
}

func handleGoogleLogin(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}

func handleGoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code for token:", err)
		return c.String(http.StatusInternalServerError, "Error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token.AccessToken,
		"token_type":   token.TokenType,
	})
}
