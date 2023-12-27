package oauth2

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/ory/fosite"
)

var (
	CreateRefreshTokenSessionQuery = `INSERT INTO RefreshTokens (TokenSignature, ClientId, TokenData, TokenExpiry) VALUES (?, ?, ?, ?)`
	GetRefreshTokenSessionQuery    = `SELECT TokenData, ClientId FROM RefreshTokens WHERE TokenSignature = ?`
	DeleteRefreshTokenSessionQuery = `DELETE FROM RefreshTokens WHERE TokenSignature = ?`
)

func CreateRefreshTokenSession(db *sql.DB, ctx context.Context, signature string, request fosite.Requester) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	expiry := request.GetSession().GetExpiresAt(fosite.RefreshToken)

	_, err = db.ExecContext(ctx, CreateAccessTokenSessionQuery, signature, request.GetClient().GetID(), jsonData, expiry)

	return err
}

func GetRefreshTokenSession(db *sql.DB, ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	var data []byte
	var clientID string

	row := db.QueryRowContext(ctx, GetRefreshTokenSessionQuery, signature)
	err := row.Scan(&data, &clientID)
	if err == sql.ErrNoRows {
		return nil, errors.New("refresh token not found")
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

func DeleteRefreshTokenSession(db *sql.DB, ctx context.Context, signature string) error {
	_, err := db.ExecContext(ctx, DeleteRefreshTokenSessionQuery, signature)
	return err
}
