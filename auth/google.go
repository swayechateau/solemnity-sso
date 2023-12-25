package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sso/auth/config"
	"sso/auth/google"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  config.GetGoogleConfig().RedirectUrl,
	ClientID:     config.GetGoogleConfig().ClientId,
	ClientSecret: config.GetGoogleConfig().ClientSecret,
	Scopes:       []string{google.Scopes.Profile, google.Scopes.Email},
	Endpoint:     google.Endpoint,
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getGoogleUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getGoogleUserInfo(state string, code string) ([]byte, error) {
	if state != randomStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	// Create a new HTTP request to the Google API
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
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

func GoogleHandler(r *mux.Router) *mux.Router {
	r.HandleFunc("/auth/google", GoogleLoginHandler)
	r.HandleFunc("/auth/google/callback", GoogleCallbackHandler)
	return r
}
