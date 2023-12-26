package models

type OAuthProvider struct {
	Id           string `db:"Id"`
	ProviderName string `db:"ProviderName"`
	ProviderId   string `db:"ProviderId"`
	Principal    string `db:"Principal"`
	Token        string `db:"Token"`
	UserId       []byte `db:"UserId"`
}

type ProviderInfo struct {
	Name string
	Id   string
}
