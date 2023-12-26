package models

import "github.com/google/uuid"

type User struct {
	Id               []byte `db:"Id"`
	Verified         bool   `db:"Verified"`
	DisplayName      string `db:"DisplayName"`
	PrimaryEmail     string `db:"PrimaryEmail"`
	PrimaryPictureId string `db:"PrimaryPictureId"`
	PrimaryLanguage  string `db:"PrimaryLanguage"`
}

type UserPicture struct {
	Id          string `db:"Id"`
	PictureType string `db:"PictureType"`
	PictureUrl  string `db:"PictureUrl"`
	UserId      []byte `db:"UserId"`
}

type UserEmail struct {
	Email     string `db:"Email"`
	IsPrimary bool   `db:"IsPrimary"`
	Verified  bool   `db:"Verified"`
	UserId    []byte `db:"UserId"`
}

func CreateUser() User {
	return User{
		Id: NewUUID(),
	}
}

func CreateUserPicture() UserPicture {
	return UserPicture{
		Id: string(NewUUID()),
	}
}

func (u *User) IdToUUID() uuid.UUID {
	return uuid.UUID(u.Id)
}
