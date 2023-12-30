package oauth2

import (
	"context"
	"database/sql"

	"github.com/ory/fosite"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{
		db: db,
	}
}

// Client
func (s *MySQLStore) CreateClient(ctx context.Context, client fosite.Client) error {
	return CreateClient(s.db, ctx, client)
}

func (s *MySQLStore) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	return GetClient(s.db, ctx, id)
}

func (s *MySQLStore) DeleteClient(ctx context.Context, id string) error {
	return DeleteClient(s.db, ctx, id)
}

// Access Token
func (s *MySQLStore) CreateAccessTokenSession(ctx context.Context, signature string, requester fosite.Requester) error {
	return CreateAccessTokenSession(s.db, ctx, signature, requester)
}

func (s *MySQLStore) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	return GetAccessTokenSession(s.db, ctx, signature, session)
}

func (s *MySQLStore) DeleteAccessTokenSession(ctx context.Context, signature string) error {
	return DeleteAccessTokenSession(s.db, ctx, signature)
}

// Refresh Token
func (s *MySQLStore) CreateRefreshTokenSession(ctx context.Context, signature string, requester fosite.Requester) error {
	return CreateRefreshTokenSession(s.db, ctx, signature, requester)
}

func (s *MySQLStore) GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	return GetRefreshTokenSession(s.db, ctx, signature, session)
}

func (s *MySQLStore) DeleteRefreshTokenSession(ctx context.Context, signature string) error {
	return DeleteRefreshTokenSession(s.db, ctx, signature)
}

// Authorization Code
func (s *MySQLStore) CreateAuthorizeCodeSession(ctx context.Context, signature string, requester fosite.Requester) error {
	return CreateAuthorizeCodeSession(s.db, ctx, signature, requester)
}

func (s *MySQLStore) GetAuthorizeCodeSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	return GetAuthorizeCodeSession(s.db, ctx, signature, session)
}

func (s *MySQLStore) DeleteAuthorizeCodeSession(ctx context.Context, signature string) error {
	return DeleteAuthorizeCodeSession(s.db, ctx, signature)
}
