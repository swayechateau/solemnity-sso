package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sso/internal/config"
	"sso/pkg/oauth2/google"

	"github.com/swayedev/way"

	"golang.org/x/oauth2"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  config.GetGoogleConfig().RedirectUrl,
	ClientID:     config.GetGoogleConfig().ClientId,
	ClientSecret: config.GetGoogleConfig().ClientSecret,
	Scopes:       []string{google.Scopes.Profile, google.Scopes.Email},
	Endpoint:     google.Endpoint,
}

var googleStateString string

func GoogleLoginHandler(c *way.Context) {
	googleStateString = randomCode()
	url := googleOauthConfig.AuthCodeURL(googleStateString)
	c.Redirect(url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(c *way.Context) {
	state := c.Request.URL.Query().Get("state")
	code := c.Request.URL.Query().Get("code")
	ctx := c.Request.Context()

	// Set the Content-Type header and write the error message
	c.Response.Header().Set("Content-Type", "application/json")

	if state != googleStateString {

		// Create a JSON response with the error message
		errMsg := fmt.Sprintf(`{"message": "%s"}`, "message: invalid oauth state")
		// c.JSON(http.StatusBadRequest, )
		http.Error(c.Response, errMsg, http.StatusBadRequest)
		return
	}

	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		// Format the error message
		er := fmt.Sprintf("code exchange failed: %s", err.Error())
		errMsg := fmt.Sprintf(`{"message": "%s"}`, er)

		http.Error(c.Response, errMsg, http.StatusBadRequest)
		return
	}

	if token != nil {
		log.Printf("Token: %v", token)
		// return
	}

	content, err := getUserInfo(c.Response, token)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		c.Redirect("/", http.StatusTemporaryRedirect)
		return
	}
	log.Printf("Content: %v", content)

	c.Response.WriteHeader(http.StatusOK)
	response := map[string]string{"Content": string(content)}

	er := json.NewEncoder(c.Response).Encode(response)
	if er != nil {
		// If encoding fails, send an internal server error
		http.Error(c.Response, "Failed to write response", http.StatusInternalServerError)
	}
}

func getUserInfo(w http.ResponseWriter, token *oauth2.Token) ([]byte, error) {
	// Use the token to fetch user info
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Now you have the user info in userInfo. You can marshal it and write it to the response.
	jsonResponse, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, "Failed to marshal user info: "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return jsonResponse, nil
}
