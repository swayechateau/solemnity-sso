package models

import (
	"sso/internal/config"
	errs "sso/internal/database/errors"
	"sso/pkg/crypt"
	"sso/pkg/database/models"
)

type Provider models.Provider

type ProviderJson struct {
	Name      string `json:"name"`
	Principal string `json:"principal"`
}

func (p *Provider) SetProviderUserId(id string) error {
	encryptedId, err := crypt.Encrypt([]byte(id), config.GetCipherKey())
	if err != nil {
		return err
	}
	p.ProviderUserIdHash = crypt.HashStringToString(id)
	p.ProviderUserId = encryptedId
	return nil
}

func (p *Provider) GetProviderUserId() (string, error) {
	decryptedId, err := crypt.Decrypt(p.ProviderUserId, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decryptedId), nil
}

func (p *Provider) SetPrincipal(principal string) error {
	encryptedPrincipal, err := crypt.Encrypt([]byte(principal), config.GetCipherKey())
	if err != nil {
		return err
	}
	p.Principal = encryptedPrincipal
	return nil
}

func (p *Provider) GetPrincipal() (string, error) {
	decryptedPrincipal, err := crypt.Decrypt(p.Principal, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decryptedPrincipal), nil
}

func (p *Provider) ToJson() ProviderJson {
	principal, _ := p.GetPrincipal()
	return ProviderJson{
		Name:      p.Name,
		Principal: principal,
	}
}

func (p *Provider) Validate() error {
	if p.Name == "" {
		return errs.ErrProviderNameEmpty
	}
	if p.ProviderUserId == "" {
		return errs.ErrProviderUserIdEmpty
	}
	if p.ProviderUserIdHash == "" {
		return errs.ErrProviderUserIdHashEmpty
	}
	if p.Principal == "" {
		return errs.ErrProviderPrincipalEmpty
	}
	return nil
}
