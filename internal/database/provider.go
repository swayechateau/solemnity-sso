package database

import (
	"sso/internal/database/models"
	"sso/pkg/database/query"

	"github.com/swayedev/way"
)

func CreateOAuthProvider(w way.Context, op *models.Provider) error {
	ctx := w.Request.Context()
	return w.PgxExecNoResult(ctx, query.CreateProvider, op.Id, op.Name, op.ProviderUserId, op.Principal, op.UserId)
}

func DeleteOAuthProvider(w way.Context, id int) error {
	ctx := w.Request.Context()
	return w.PgxExecNoResult(ctx, query.DeleteProviderById, id)
}
