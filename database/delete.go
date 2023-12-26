package database

import (
	"context"
	"sso/database/query"
)

func (conn *Conn) DeleteUser(ctx context.Context, id []byte) error {
	return conn.Delete(ctx, query.DeleteUserById, id)
}

func (conn *Conn) DeleteUserPicture(ctx context.Context, id string) error {
	return conn.Delete(ctx, query.DeleteUserPictureById, id)
}

func (conn *Conn) DeleteUserEmail(ctx context.Context, email string) error {
	return conn.Delete(ctx, query.DeleteUserEmailByEmail, email)
}

func (conn *Conn) DeleteOAuthProvider(ctx context.Context, providerName string, providerId string) error {
	return conn.Delete(ctx, query.DeleteOAuthProviderById, providerName, providerId)
}

func (conn *Conn) DeleteClient(ctx context.Context, id string) error {
	return conn.Delete(ctx, query.DeleteClientById, id)
}

func (conn *Conn) DeleteAccessToken(ctx context.Context, signature string) error {
	return conn.Delete(ctx, query.DeleteAccessTokenBySig, signature)
}

func (conn *Conn) DeleteRefreshToken(ctx context.Context, signature string) error {
	return conn.Delete(ctx, query.DeleteRefreshTokenBySig, signature)
}

func (conn *Conn) DeleteAuthCode(ctx context.Context, signature string) error {
	return conn.Delete(ctx, query.DeleteAuthCodeBySig, signature)
}

func (conn *Conn) DeleteUserConsent(ctx context.Context, id string) error {
	return conn.Delete(ctx, query.DeleteUserConsentById, id)
}

func (conn *Conn) DeleteScope(ctx context.Context, id string) error {
	return conn.Delete(ctx, query.DeleteScopeById, id)
}
