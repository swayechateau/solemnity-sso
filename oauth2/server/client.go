package oauth2

import (
	"context"
	"database/sql"

	"github.com/ory/fosite"
	"github.com/pkg/errors"
)

var (
	CreateClientQuery = "INSERT INTO Clients (Id, ClientSecret, RedirectUri, Scopes, GrantTypes) VALUES (?, ?, ?, ?, ?)"
	GetClientQuery    = "SELECT Id, ClientSecret, RedirectUri, Scopes, GrantTypes FROM Clients WHERE Id = ?"
	DeleteClientQuery = "DELETE FROM Clients WHERE Id = ?"
)

func CreateClient(db *sql.DB, ctx context.Context, client fosite.Client) error {
	_, err := db.ExecContext(ctx, CreateClientQuery, client.GetID(), client.GetHashedSecret(), client.GetRedirectURIs(), client.GetScopes(), client.GetGrantTypes())
	return err
}

func GetClient(db *sql.DB, ctx context.Context, id string) (fosite.Client, error) {
	var client fosite.DefaultClient

	row := db.QueryRowContext(ctx, GetClientQuery, id)
	err := row.Scan(&client.ID, &client.Secret, &client.RedirectURIs, &client.Scopes, &client.GrantTypes)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(fosite.ErrNotFound, "client not found")
		}
		return nil, err
	}

	return &client, nil
}

func DeleteClient(db *sql.DB, ctx context.Context, id string) error {
	_, err := db.ExecContext(ctx, DeleteClientQuery, id)
	return err
}
