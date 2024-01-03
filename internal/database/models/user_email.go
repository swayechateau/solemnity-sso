package models

import (
	"sso/internal/config"
	"sso/pkg/crypt"
	"sso/pkg/database/models"
)

type UserEmail models.UserEmail
type UserEmailJson struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

func (e *UserEmail) SetEmail(email string) error {
	encrypted, err := crypt.Encrypt([]byte(email), config.GetCipherKey())
	if err != nil {
		return err
	}
	e.EmailHash = crypt.HashStringToString(email)
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
