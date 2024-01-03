package models

import "time"

type Provider struct {
	Id                 int       `json:"-" db:"Id"`
	Name               string    `json:"provider_name" db:"ProviderName"`
	Principal          string    `json:"principal" db:"Principal"`
	ProviderUserId     string    `json:"provider_user_id" db:"ProviderId"`
	ProviderUserIdHash string    `json:"-" db:"ProviderIdHash"`
	UserId             [16]byte  `json:"-" db:"UserId"`
	CreatedAt          time.Time `json:"created_at" db:"CreatedAt"`
	UpdatedAt          time.Time `json:"updated_at" db:"UpdatedAt"`
}
