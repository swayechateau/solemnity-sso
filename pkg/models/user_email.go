package models

type UserEmail struct {
	Email    string   `json:"email" db:"Email"`
	Primary  bool     `json:"primary" db:"IsPrimary"`
	Verified bool     `json:"verified" db:"Verified"`
	UserId   [16]byte `json:"-" db:"UserId"`
}
