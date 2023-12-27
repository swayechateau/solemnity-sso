package provider

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthContent struct {
	Token   *oauth2.Token
	Content []byte
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

func GetOAuthInfoWithToken(token *oauth2.Token, api string) (*OAuthContent, error) {
	contents, err := GetOAuthInfo(token, api)
	if err != nil {
		return nil, err
	}

	return &OAuthContent{
		Token:   token,
		Content: contents,
	}, nil
}
