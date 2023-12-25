package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sso/auth/config"
	"sso/auth/microsoft"

	"github.com/gorilla/mux"
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

func MicrosoftLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := microsoftOauthConfig.AuthCodeURL(randomStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func MicrosoftCallbackHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getMicrosoftUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getMicrosoftUserInfo(state string, code string) ([]byte, error) {
	if state != randomStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := microsoftOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	// Create a new HTTP request to the Microsoft Graph API
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err.Error())
	}

	// Add the access token in the Authorization header
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	// Create an HTTP client and send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	// Read and return the response body
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func MicrosoftHandler(r *mux.Router) *mux.Router {
	r.HandleFunc("/auth/microsoft", MicrosoftLoginHandler)
	r.HandleFunc("/auth/microsoft/callback", MicrosoftCallbackHandler)
	return r
}
