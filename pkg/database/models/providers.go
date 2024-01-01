package models

type Provider struct {
	Id             int      `json:"-" db:"Id"`
	Name           string   `json:"provider_name" db:"ProviderName"`
	Principal      string   `json:"principal" db:"Principal"`
	ProviderUserId string   `json:"provider_user_id" db:"ProviderId"`
	UserId         [16]byte `json:"-" db:"UserId"`
}
