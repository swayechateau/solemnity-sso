package models

import (
	"sso/internal/config"
	"sso/pkg/crypt"
	"sso/pkg/database/models"
)

type User models.User
type UserEmail models.UserEmail
type UserPicture models.UserPicture

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
