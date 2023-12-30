package models

type Provider struct {
	Id             string `json:"id" db:"Id"`
	Name           string `json:"provider_name" db:"ProviderName"`
	Principal      string `json:"principal" db:"Principal"`
	ProviderUserId string `json:"provider_user_id" db:"ProviderId"`
	UserId         []byte `json:"-" db:"UserId"`
}
