package models

import (
	"time"

	"github.com/google/uuid"
)

type UserPicture struct {
	Id        [16]byte  `json:"id" db:"Id"`
	Extension string    `json:"type" db:"Extension"` // gif, jpeg, png
	Uri       string    `json:"url" db:"Uri"`
	UserId    [16]byte  `json:"-" db:"UserId"`
	CreatedAt time.Time `json:"created_at" db:"CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" db:"UpdatedAt"`
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
