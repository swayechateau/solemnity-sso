package user

import (
	"sso/database/models"

	"github.com/google/uuid"
)

type AuthUser struct {
	Id              uuid.UUID       `json:"id"`
	Verified        bool            `json:"verified"`
	DisplayName     string          `json:"display_name"`
	PrimaryEmail    string          `json:"primary_email"`
	PrimaryPicture  string          `json:"primary_picture"`
	PrimaryLanguage string          `json:"primary_language"`
	Pictures        []UserPicture   `json:"profile_pictures"`
	Email           []AuthUserEmail `json:"email"`
	OAuthProviders  []OAuthProvider `json:"oauth_providers"`
}

type UserPicture struct {
	Id   string `json:"id"`
	Type string `json:"type"` // gif, jpeg, png
	Url  string `json:"url"`
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
	// Token     *oauth2.Token `json:"token"`
}

func (u *AuthUser) IdToByte() []byte {
	return u.Id[:]
}

func (u *AuthUser) IdToString() string {
	return u.Id.String()
}

func (u *AuthUser) ToUser() *models.User {
	return &models.User{
		Id:               u.IdToByte(),
		Verified:         u.Verified,
		DisplayName:      u.DisplayName,
		PrimaryEmail:     u.PrimaryEmail,
		PrimaryPictureId: u.PrimaryPicture,
		PrimaryLanguage:  u.PrimaryLanguage,
	}
}
