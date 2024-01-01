package models

import "github.com/google/uuid"

type UserPicture struct {
	Id     [16]byte `json:"id" db:"Id"`
	Type   string   `json:"type" db:"Type"` // gif, jpeg, png
	Url    string   `json:"url" db:"Url"`
	UserId [16]byte `json:"-" db:"UserId"`
}

// UserPicture ID functions
func (up *UserPicture) NewId() {
	up.Id = uuid.New()
}

func (up *UserPicture) IdToUUID() uuid.UUID {
	return uuid.UUID(up.Id)
}

func (up *UserPicture) IdToString() string {
	return up.IdToUUID().String()
}

func (up *UserPicture) SetIdFromString(id string) {
	up.Id = uuid.MustParse(id)
}

func (up *UserPicture) SetIdFromUUID(id uuid.UUID) {
	up.Id = id
}
