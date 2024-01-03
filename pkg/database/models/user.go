package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               [16]byte      `json:"id" db:"Id"`
	Verified         bool          `json:"verified" db:"Verified"`
	DisplayName      string        `json:"display_name" db:"DisplayName"`
	PrimaryEmailHash string        `json:"-" db:"PrimaryEmailHash"`
	PrimaryEmail     string        `json:"primary_email" db:"PrimaryEmailAddress"`
	PrimaryPictureId [16]byte      `json:"primary_picture" db:"PrimaryPicture"`
	PrimaryLanguage  string        `json:"primary_language" db:"PrimaryLanguage"`
	Pictures         []UserPicture `json:"pictures" db:"-"`
	Email            []UserEmail   `json:"emails" db:"-"`
	Providers        []Provider    `json:"providers" db:"-"`
	CreatedAt        time.Time     `json:"created_at" db:"CreatedAt"`
	UpdatedAt        time.Time     `json:"updated_at" db:"UpdatedAt"`
}

// User ID functions
func (u *User) NewId() {
	u.Id, _ = uuid.NewRandom()
}

func (u *User) IdToUUID() uuid.UUID {
	return uuid.UUID(u.Id)
}

func (u *User) IdToString() string {
	return u.IdToUUID().String()
}

func (u *User) SetIdFromString(id string) {
	u.Id = uuid.MustParse(id)
}

func (u *User) SetIdFromUUID(id uuid.UUID) {
	u.Id = id
}

func (u *User) PrimaryPictureIdToUUID() uuid.UUID {
	return uuid.UUID(u.Id)
}

func (u *User) PrimaryPictureIdToString() string {
	return u.IdToUUID().String()
}

func (u *User) SetPrimaryPictureIdFromString(id string) {
	u.Id = uuid.MustParse(id)
}

func (u *User) SetPrimaryPictureIdFromUUID(id uuid.UUID) {
	u.Id = id
}
