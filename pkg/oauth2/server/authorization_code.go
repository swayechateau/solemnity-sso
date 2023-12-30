package oauth2

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/ory/fosite"
)

var (
	DeleteAuthorizeCodeSessionQuery = `DELETE FROM AuthCodes WHERE CodeSignature = ?`
	CreateAuthorizeCodeSessionQuery = `INSERT INTO AuthCodes (CodeSignature, ClientId, CodeData, CodeExpiry) VALUES (?, ?, ?, ?)`
	GetAuthorizeCodeSessionQuery    = `SELECT CodeData, ClientId FROM AuthCodes WHERE CodeSignature = ?`
)

func CreateAuthorizeCodeSession(db *sql.DB, ctx context.Context, code string, request fosite.Requester) error {
	// Serialize the request data to store in the database
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	expiry := request.GetSession().GetExpiresAt(fosite.AuthorizeCode)

	// Prepare SQL query to insert the authorization code session
	_, err = db.ExecContext(ctx, CreateAuthorizeCodeSessionQuery, code, request.GetClient().GetID(), jsonData, expiry)
	return err
}

func GetAuthorizeCodeSession(db *sql.DB, ctx context.Context, code string, session fosite.Session) (fosite.Requester, error) {
	var data []byte
	var clientID string

	row := db.QueryRowContext(ctx, GetAuthorizeCodeSessionQuery, code)
	err := row.Scan(&data, &clientID)
	if err == sql.ErrNoRows {
		return nil, errors.New("authorization code not found")
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &session)
	if err != nil {
		return nil, err
	}

	requester := fosite.NewRequest()
	requester.SetSession(session)

	// Here, you need to set the client to the requester using your own method, as fosite doesn't expose SetClient

	return requester, nil
}

func DeleteAuthorizeCodeSession(db *sql.DB, ctx context.Context, signature string) error {
	_, err := db.ExecContext(ctx, DeleteAuthorizeCodeSessionQuery, signature)
	return err
}
