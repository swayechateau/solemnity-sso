package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"sso/auth/config"
	"sso/auth/github"

	"github.com/gorilla/mux"
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

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL(randomStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getGithubUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getGithubUserInfo(state string, code string) ([]byte, error) {
	if state != randomStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err.Error())
	}

	req.Header.Add("Authorization", "token "+token.AccessToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func GithubHandler(r *mux.Router) *mux.Router {
	r.HandleFunc("/auth/github", GithubLoginHandler)
	r.HandleFunc("/auth/github/callback", GithubCallbackHandler)
	return r
}
