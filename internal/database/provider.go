package database

import (
	"sso/internal/database/models"
	"sso/pkg/database/query"

	"github.com/swayedev/way"
)

// Find Providers by UserId
func FindUserProvidersByUserId(w *way.Context, userId [16]byte) ([]*models.Provider, error) {
	var userProviders []*models.Provider
	ctx := w.Request.Context()

	rows, err := w.PgxQuery(ctx, query.FindProviderByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var up models.Provider
		if err := rows.Scan(&up.Id, &up.Name, &up.ProviderUserId, &up.ProviderUserIdHash, &up.Principal, &up.UserId, &up.CreatedAt, &up.UpdatedAt); err != nil {
			return nil, err
		}
		userProviders = append(userProviders, &up)
	}

	return userProviders, nil
}

func DeleteOAuthProvider(w *way.Context, id int) error {
	ctx := w.Request.Context()
	return w.PgxExecNoResult(ctx, query.DeleteProviderById, id)
}

func CreateProvider(w *way.Context, p models.Provider) error {
	ctx := w.Request.Context()
	if err := p.Validate(); err != nil {
		return err
	}
	return w.PgxExecNoResult(ctx, query.CreateProvider, p.Name, p.ProviderUserId, p.ProviderUserIdHash, p.Principal, p.UserId)
}
