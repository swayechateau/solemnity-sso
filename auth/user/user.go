package user

import "encoding/json"

type AuthUser struct {
	Id             string          `json:"id"`
	Verified       bool            `json:"verified"`
	Email          []AuthUserEmail `json:"email"`
	ProfilePic     string          `json:"profile_picture"`
	OAuthProviders []OAuthProvider `json:"oauth_providers"`
}

type AuthUserEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

type OAuthProvider struct {
	Id        string `json:"id"`
	Provider  string `json:"provider"`
	Principal string `json:"principal"`
	// Context  string `json:"context"`
	// AccessToken  string `json:"access_token"`
	// RefreshToken string `json:"refresh_token"`
	// Expiry       int64  `json:"expiry"`
}

func SendUser(data interface{}) (d []byte) {
	d, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return d
}
