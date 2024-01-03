package models

import (
	"sso/internal/config"
	errs "sso/internal/database/errors"
	"sso/pkg/crypt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               [16]byte      `json:"id" db:"Id"`
	Verified         bool          `json:"verified" db:"Verified"`
	DisplayName      string        `json:"display_name" db:"DisplayName"`
	PrimaryEmail     string        `json:"primary_email" db:"PrimaryEmail"`
	PrimaryEmailHash string        `json:"-" db:"PrimaryEmailHash"`
	PrimaryPictureId [16]byte      `json:"primary_picture" db:"PrimaryPicture"`
	PrimaryLanguage  string        `json:"primary_language" db:"PrimaryLanguage"`
	Pictures         []UserPicture `json:"pictures,omitempty" db:"-"`
	Emails           []UserEmail   `json:"emails,omitempty" db:"-"`
	Providers        []Provider    `json:"providers,omitempty" db:"-"`
	CreatedAt        time.Time     `json:"created_at" db:"CreatedAt"`
	UpdatedAt        time.Time     `json:"updated_at" db:"UpdatedAt"`
}

type UserJson struct {
	Id               string            `json:"id"`
	Verified         bool              `json:"verified"`
	DisplayName      string            `json:"display_name"`
	PrimaryEmail     string            `json:"primary_email"`
	PrimaryPictureId string            `json:"primary_picture_id"`
	PrimaryLanguage  string            `json:"primary_language"`
	Emails           []UserEmailJson   `json:"emails,omitempty"`
	Pictures         []UserPictureJson `json:"pictures,omitempty"`
	Providers        []ProviderJson    `json:"providers,omitempty"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

func NewUser() User {
	u := User{}
	u.Id = uuid.New()
	u.PrimaryLanguage = "en"
	return u
}
func (u *User) GetUuid() uuid.UUID {
	return u.Id
}

func (u *User) SetIdFromString(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	u.Id = uuid
	return nil
}

func (u *User) SetIdFromBytes(id []byte) {
	u.Id = uuid.UUID(id)
}

func (u *User) GetPrimaryPictureUuid() uuid.UUID {
	return u.PrimaryPictureId
}

func (u *User) SetPrimaryPictureIdFromBytes(id []byte) {
	u.PrimaryPictureId = uuid.UUID(id)
}

func (u *User) SetPrimaryPictureIdFromString(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	u.PrimaryPictureId = uuid
	return nil
}

func (u *User) SetDisplayName(displayName string) error {
	encrypted, err := crypt.Encrypt([]byte(displayName), config.GetCipherKey())
	if err != nil {
		return err
	}
	u.DisplayName = encrypted
	return nil
}

func (u *User) GetDisplayName() (string, error) {
	decrypted, err := crypt.Decrypt(u.DisplayName, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (u *User) SetPrimaryEmail(email string) error {
	encrypted, err := crypt.Encrypt([]byte(email), config.GetCipherKey())
	if err != nil {
		return err
	}
	u.PrimaryEmailHash = crypt.HashStringToString(email)
	u.PrimaryEmail = encrypted
	return nil
}

func (u *User) GetPrimaryEmail() (string, error) {
	decrypted, err := crypt.Decrypt(u.PrimaryEmail, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (u *User) ToJson() UserJson {
	var userJson UserJson
	userJson.Id = u.GetUuid().String()
	userJson.Verified = u.Verified
	userJson.PrimaryLanguage = u.PrimaryLanguage
	userJson.DisplayName, _ = u.GetDisplayName()
	userJson.PrimaryEmail, _ = u.GetPrimaryEmail()
	userJson.PrimaryPictureId = u.GetPrimaryPictureUuid().String()
	userJson.CreatedAt = u.CreatedAt
	userJson.UpdatedAt = u.UpdatedAt

	for _, picture := range u.Pictures {
		userJson.Pictures = append(userJson.Pictures, picture.ToJson())
	}

	for _, email := range u.Emails {
		userJson.Emails = append(userJson.Emails, email.ToJson())
	}

	for _, provider := range u.Providers {
		userJson.Providers = append(userJson.Providers, provider.ToJson())
	}

	return userJson
}

func (u *User) Validate() error {
	if u.Id == uuid.Nil {
		return errs.ErrIdEmpty
	}
	if u.PrimaryEmail == "" {
		return errs.ErrEmailEmpty
	}
	if u.PrimaryEmailHash == "" {
		return errs.ErrEmailHashEmpty
	}

	return nil
}
