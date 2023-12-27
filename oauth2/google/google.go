package google

import (
	"encoding/json"
	"sso/oauth2/scopes"

	"golang.org/x/oauth2"
)

var (
	Endpoint = oauth2.Endpoint{
		AuthURL:   "https://accounts.google.com/o/oauth2/auth",
		TokenURL:  "https://oauth2.googleapis.com/token",
		AuthStyle: oauth2.AuthStyleInParams,
	}
	Scopes = scopes.OAuthScopes{
		Profile: "https://www.googleapis.com/auth/userinfo.profile",
		Email:   "https://www.googleapis.com/auth/userinfo.email",
	}
	Api = "https://www.googleapis.com/oauth2/v3/userinfo" //"https://www.googleapis.com/oauth2/v2/userinfo"
)

type GoogleContext struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}

func JsonToContext(content []byte) GoogleContext {
	var googleContext GoogleContext
	err := json.Unmarshal(content, &googleContext)
	if err != nil {
		return GoogleContext{}
	}
	return googleContext
}

func (googleContext *GoogleContext) ToJson() []byte {
	json, err := json.Marshal(googleContext)
	if err != nil {
		return []byte{}
	}
	return json
}
