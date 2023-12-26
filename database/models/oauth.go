package models

// OAuth2 Server Tables
type Client struct {
	Id           string `db:"Id"`
	ClientSecret string `db:"ClientSecret"`
	RedirectUri  string `db:"RedirectUri"`
	Scopes       string `db:"Scopes"`
	GrantTypes   string `db:"GrantTypes"`
}

type AccessToken struct {
	TokenSignature string `db:"TokenSignature"`
	ClientId       string `db:"ClientId"`
	TokenData      []byte `db:"TokenData"`
	TokenExpiry    string `db:"TokenExpiry"`
}

type RefreshToken struct {
	TokenSignature string `db:"TokenSignature"`
	ClientId       string `db:"ClientId"`
	TokenData      []byte `db:"TokenData"`
	TokenExpiry    string `db:"TokenExpiry"`
}

type AuthCode struct {
	CodeSignature string `db:"CodeSignature"`
	ClientId      string `db:"ClientId"`
	CodeData      []byte `db:"CodeData"`
	CodeExpiry    string `db:"CodeExpiry"`
}

type UserConsent struct {
	Id       int    `db:"Id"`
	UserId   []byte `db:"UserId"`
	ClientId string `db:"ClientId"`
	Scopes   string `db:"Scopes"`
}

type Scope struct {
	Id               int    `db:"Id"`
	ScopeName        string `db:"ScopeName"`
	ScopeDescription string `db:"ScopeDescription"`
}
