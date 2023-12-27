package github

import (
	"encoding/json"
	"sso/oauth2/scopes"

	"golang.org/x/oauth2"
)

// Endpoint is Github's OAuth 2.0 endpoint.
var (
	Endpoint = oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	}
	Scopes = scopes.OAuthScopes{
		Profile: "user",
		Email:   "user:email",
	}
	Api = "https://api.github.com/user"
)

type GithubContext struct {
	Login                   string `json:"login"`
	Id                      int64  `json:"id"`
	NodeId                  string `json:"node_id"`
	AvatarUrl               string `json:"avatar_url"`
	GravatarId              string `json:"gravatar_id"`
	Url                     string `json:"url"`
	HtmlUrl                 string `json:"html_url"`
	FollowersUrl            string `json:"followers_url"`
	FollowingUrl            string `json:"following_url"`
	GistsUrl                string `json:"gists_url"`
	StarredUrl              string `json:"starred_url"`
	SubscriptionsUrl        string `json:"subscriptions_url"`
	OrganizationsUrl        string `json:"organizations_url"`
	ReposUrl                string `json:"repos_url"`
	EventsUrl               string `json:"events_url"`
	ReceivedEventsUrl       string `json:"received_events_url"`
	Type                    string `json:"type"`
	SiteAdmin               bool   `json:"site_admin"`
	Name                    string `json:"name"`
	Company                 string `json:"company"`
	Blog                    string `json:"blog"`
	Location                string `json:"location"`
	Email                   string `json:"email"`
	Hireable                string `json:"hireable"`
	Bio                     string `json:"bio"`
	TwitterUsername         string `json:"twitter_username"`
	PublicRepos             int64  `json:"public_repos"`
	PublicGists             int64  `json:"public_gists"`
	Followers               int64  `json:"followers"`
	Following               int64  `json:"following"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
	PrivateGists            int64  `json:"private_gists"`
	TotalPrivateRepos       int64  `json:"total_private_repos"`
	OwnedPrivateRepos       int64  `json:"owned_private_repos"`
	DiskUsage               int64  `json:"disk_usage"`
	Collaborators           int64  `json:"collaborators"`
	TwoFactorAuthentication bool   `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int64  `json:"space"`
		Collaborators int64  `json:"collaborators"`
		PrivateRepos  int64  `json:"private_repos"`
	} `json:"plan"`
}

func JsonToContext(content []byte) GithubContext {
	var githubContext GithubContext
	err := json.Unmarshal(content, &githubContext)
	if err != nil {
		return GithubContext{}
	}
	return githubContext
}

func (githubContext *GithubContext) ToJson() []byte {
	json, err := json.Marshal(githubContext)
	if err != nil {
		return []byte{}
	}
	return json
}
