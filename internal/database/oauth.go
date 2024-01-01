package database

import (
	"sso/pkg/database/models"
	"sso/pkg/database/query"

	"github.com/swayedev/way"
)

func CreateClient(w way.Context, c *models.Client) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateClient, c)
}

func DeleteClient(w way.Context, id string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteClientById, id)
}

func CreateAccessToken(w way.Context, at *models.AccessToken) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateAccessToken, at)
}

func DeleteAccessToken(w way.Context, signature string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteAccessTokenBySig, signature)
}

func CreateRefreshToken(w way.Context, rt *models.RefreshToken) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateRefreshToken, rt)
}

func DeleteRefreshToken(w way.Context, signature string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteRefreshTokenBySig, signature)
}

func CreateAuthCode(w way.Context, ac *models.AuthCode) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateAuthCode, ac)
}

func DeleteAuthCode(w way.Context, signature string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteAuthCodeBySig, signature)
}

func CreateUserConsent(w way.Context, uc *models.UserConsent) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateUserConsent, uc)
}

func DeleteUserConsent(w way.Context, id string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteUserConsentById, id)
}

func CreateScope(w way.Context, s *models.Scope) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateScope, s)
}

func DeleteScope(w way.Context, id string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteScopeById, id)
}
