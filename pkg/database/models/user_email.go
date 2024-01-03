package models

import "time"

type UserEmail struct {
	EmailHash string    `json:"-" db:"EmailHash"`
	Email     string    `json:"email" db:"EmailAddress"`
	Primary   bool      `json:"primary" db:"IsPrimary"`
	Verified  bool      `json:"verified" db:"Verified"`
	UserId    [16]byte  `json:"-" db:"UserId"`
	CreatedAt time.Time `json:"created_at" db:"CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" db:"UpdatedAt"`
}
