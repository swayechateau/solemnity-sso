package models

import (
	"sso/internal/config"
	"sso/pkg/crypt"
	"sso/pkg/database/models"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               [16]byte      `json:"id" db:"Id"`
	Verified         bool          `json:"verified" db:"Verified"`
	DisplayName      string        `json:"display_name" db:"DisplayName"`
	PrimaryEmail     string        `json:"primary_email" db:"PrimaryEmail"`
	PrimaryPictureId [16]byte      `json:"primary_picture" db:"PrimaryPicture"`
	PrimaryLanguage  string        `json:"primary_language" db:"PrimaryLanguage"`
	Pictures         []UserPicture `json:"pictures,omitempty" db:"-"`
	Email            []UserEmail   `json:"emails,omitempty" db:"-"`
	Providers        []Provider    `json:"providers,omitempty" db:"-"`
	CreatedAt        time.Time     `json:"created_at" db:"CreatedAt"`
	UpdatedAt        time.Time     `json:"updated_at" db:"UpdatedAt"`
}

type UserEmail models.UserEmail
type UserPicture models.UserPicture

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

type UserEmailJson struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

type UserPictureJson struct {
	Id        string `json:"id"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
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
	userJson.Emails = make([]UserEmailJson, len(u.Email))
	userJson.Pictures = make([]UserPictureJson, len(u.Pictures))
	userJson.Providers = make([]ProviderJson, len(u.Providers))
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

	for _, email := range u.Email {
		userJson.Emails = append(userJson.Emails, email.ToJson())
	}

	for _, provider := range u.Providers {
		userJson.Providers = append(userJson.Providers, provider.ToJson())
	}

	return userJson
}

func (e *UserEmail) SetEmail(email string) error {
	encrypted, err := crypt.Encrypt([]byte(email), config.GetCipherKey())
	if err != nil {
		return err
	}
	e.Email = encrypted
	return nil
}

func (e *UserEmail) GetEmail() (string, error) {
	decrypted, err := crypt.Decrypt(e.Email, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (e *UserEmail) ToJson() UserEmailJson {
	var userEmailJson UserEmailJson
	userEmailJson.Email, _ = e.GetEmail()
	userEmailJson.Primary = e.Primary
	userEmailJson.Verified = e.Verified
	return userEmailJson
}

func (p *UserPicture) GetUuid() uuid.UUID {
	return p.Id
}

func (p *UserPicture) SetIdFromBytes(id []byte) {
	p.Id = uuid.UUID(id)
}

func (p *UserPicture) SetIdFromString(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	p.Id = uuid
	return nil
}

func (p *UserPicture) SetPictureUrl(url string) error {
	encrypted, err := crypt.Encrypt([]byte(url), config.GetCipherKey())
	if err != nil {
		return err
	}
	p.Url = encrypted
	return nil
}

func (p *UserPicture) GetPictureUrl() (string, error) {
	decrypted, err := crypt.Decrypt(p.Url, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (p *UserPicture) ToJson() UserPictureJson {
	var userPictureJson UserPictureJson
	userPictureJson.Id = p.GetUuid().String()
	userPictureJson.Url, _ = p.GetPictureUrl()
	return userPictureJson
}
