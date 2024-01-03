package models

import (
	"sso/internal/config"
	"sso/pkg/crypt"
	"sso/pkg/database/models"

	"github.com/google/uuid"
)

type UserPicture models.UserPicture
type UserPictureJson struct {
	Id        string `json:"id"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
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

func (p *UserPicture) SetUrl(url string) error {
	encrypted, err := crypt.Encrypt([]byte(url), config.GetCipherKey())
	if err != nil {
		return err
	}
	p.Uri = encrypted
	return nil
}

func (p *UserPicture) GetUrl() (string, error) {
	decrypted, err := crypt.Decrypt(p.Uri, config.GetCipherKey())
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (p *UserPicture) SetExtension(extension string) {
	p.Extension = extension
}

func (p *UserPicture) ToJson() UserPictureJson {
	var userPictureJson UserPictureJson
	userPictureJson.Id = p.GetUuid().String()
	userPictureJson.Extension = p.Extension
	userPictureJson.Url, _ = p.GetUrl()
	return userPictureJson
}
