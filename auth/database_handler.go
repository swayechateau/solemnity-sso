package auth

import "sso/database"

type AuthDB struct {
	Auth *database.Conn
}

func NewAuthDBHandler(app *database.Conn) *AuthDB {
	return &AuthDB{Auth: app}
}
