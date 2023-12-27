package microsoft

import (
	"encoding/json"
	"sso/oauth2/scopes"

	"golang.org/x/oauth2"
)

// LiveConnectEndpoint is Windows's Live ID OAuth 2.0 endpoint.
var (
	LiveConnectEndpoint = oauth2.Endpoint{
		AuthURL:  "https://login.live.com/oauth20_authorize.srf",
		TokenURL: "https://login.live.com/oauth20_token.srf",
	}
	Scopes = scopes.OAuthScopes{
		Profile: "User.Read",
		Email:   "User.Read",
	}
	Api = "https://graph.microsoft.com/v1.0/me"
)

func AzureADEndpoint(tenant string) oauth2.Endpoint {
	if tenant == "" {
		tenant = "common"
	}
	return oauth2.Endpoint{
		AuthURL:  "https://login.microsoftonline.com/" + tenant + "/oauth2/v2.0/authorize",
		TokenURL: "https://login.microsoftonline.com/" + tenant + "/oauth2/v2.0/token",
	}
}

type MicrosoftContext struct {
	DataContext       string   `json:"@odata.context"`
	BusinessPhones    []string `json:"businessPhones"`
	DisplayName       string   `json:"displayName"`
	GivenName         string   `json:"givenName"`
	JobTitle          string   `json:"jobTitle"`
	Mail              string   `json:"mail"`
	MobilePhone       string   `json:"mobilePhone"`
	OfficeLocation    string   `json:"officeLocation"`
	PreferredLanguage string   `json:"preferredLanguage"`
	Surname           string   `json:"surname"`
	UserPrincipalName string   `json:"userPrincipalName"`
	Id                string   `json:"id"`
}

func JsonToContext(content []byte) MicrosoftContext {
	var microsoftContext MicrosoftContext
	err := json.Unmarshal(content, &microsoftContext)
	if err != nil {
		return MicrosoftContext{}
	}
	return microsoftContext
}

func (microsoftContext *MicrosoftContext) ToJson() []byte {
	json, err := json.Marshal(microsoftContext)
	if err != nil {
		return []byte{}
	}
	return json
}
