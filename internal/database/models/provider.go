package models

import (
	"sso/internal/config"
	"sso/pkg/crypt"
	"sso/pkg/database/models"
)

type Provider models.Provider

func (p *Provider) SetProviderUserId(id string) error {
	encryptedId, err := crypt.Encrypt([]byte(id), config.GetCipherKey())
	if err != nil {
		return err
	}
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
