package user

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type AuthUser struct {
	Id              string                 `json:"id"`
	Verified        bool                   `json:"verified"`
	DisplayName     string                 `json:"display_name"`
	PrimaryEmail    string                 `json:"primary_email"`
	PrimaryPicture  string                 `json:"primary_picture"`
	PrimaryLanguage string                 `json:"primary_language"`
	Pictures        map[string]UserPicture `json:"profile_picture"`
	Email           []AuthUserEmail        `json:"email"`
	OAuthProviders  []OAuthProvider        `json:"oauth_providers"`
}

type UserPicture struct {
	Id   string `json:"id"`
	Type string `json:"type"` // gif, jpeg, png
	Url  string `json:"url"`
}

type AuthUserEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

type OAuthProvider struct {
	Id        string        `json:"id"`
	Provider  string        `json:"provider"`
	Principal string        `json:"principal"`
	Token     *oauth2.Token `json:"token"`
}

func GetOAuthInfo(token *oauth2.Token, api string) ([]byte, error) {
	req, err := http.NewRequest("GET", api, nil)
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
