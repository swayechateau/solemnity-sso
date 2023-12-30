package oauth2

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/ory/fosite"
)

var (
	CreateAccessTokenSessionQuery = `INSERT INTO AccessTokens (TokenSignature, ClientId, TokenData, TokenExpiry) VALUES (?, ?, ?, ?)`
	GetAccessTokenSessionQuery    = `SELECT TokenData, ClientId FROM AccessTokens WHERE TokenSignature = ?`
	DeleteAccessTokenSessionQuery = `DELETE FROM AccessTokens WHERE TokenSignature = ?`
)

func CreateAccessTokenSession(db *sql.DB, ctx context.Context, signature string, request fosite.Requester) error {
	// Serialize the request data to store in the database
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	expiry := request.GetSession().GetExpiresAt(fosite.AccessToken)

	// Prepare SQL query
	_, err = db.ExecContext(ctx, CreateAccessTokenSessionQuery, signature, request.GetClient().GetID(), jsonData, expiry)

	return err
}

func GetAccessTokenSession(db *sql.DB, ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	var data []byte
	var clientID string

	row := db.QueryRowContext(ctx, GetAccessTokenSessionQuery, signature)
	err := row.Scan(&data, &clientID)
	if err == sql.ErrNoRows {
		return nil, errors.New("access token not found")
	} else if err != nil {
		return nil, err
	}

	// Deserialize the session data
	err = json.Unmarshal(data, &session)
	if err != nil {
		return nil, err
	}

	// Construct and return the fosite requester object
	requester := fosite.NewRequest()
	requester.SetSession(session)

	// Here, you need to set the client to the requester using your own method, as fosite doesn't expose SetClient

	return requester, nil
}

func DeleteAccessTokenSession(db *sql.DB, ctx context.Context, signature string) error {
	_, err := db.ExecContext(ctx, DeleteAccessTokenSessionQuery, signature)
	return err
}
