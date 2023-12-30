package models

import "github.com/google/uuid"

type User struct {
	Id              [16]byte      `json:"id" db:"Id"`
	Verified        bool          `json:"verified" db:"Verified"`
	DisplayName     string        `json:"display_name" db:"DisplayName"`
	PrimaryEmail    string        `json:"primary_email" db:"PrimaryEmail"`
	PrimaryPicture  string        `json:"primary_picture" db:"PrimaryPicture"`
	PrimaryLanguage string        `json:"primary_language" db:"PrimaryLanguage"`
	Pictures        []UserPicture `json:"pictures" db:"-"`
	Email           []UserEmail   `json:"emails" db:"-"`
	Providers       []Provider    `json:"providers" db:"-"`
}

type UserPicture struct {
	Id     string   `json:"id" db:"Id"`
	Type   string   `json:"type" db:"Type"` // gif, jpeg, png
	Url    string   `json:"url" db:"Url"`
	UserId [16]byte `json:"-" db:"UserId"`
}

type UserEmail struct {
	Email    string   `json:"email" db:"Email"`
	Primary  bool     `json:"primary" db:"IsPrimary"`
	Verified bool     `json:"verified" db:"Verified"`
	UserId   [16]byte `json:"-" db:"UserId"`
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
