package database

import (
	"context"
	"database/sql"
	"sso/database/models"
	"sso/database/query"
)

func (conn *Conn) FindUserById(ctx context.Context, id []byte) (*models.User, error) {
	var u models.User

	row := conn.QueryRow(ctx, query.FindUserById, id)
	if err := row.Scan(&u.Id, &u.Verified, &u.DisplayName, &u.PrimaryEmail, &u.PrimaryPictureId, &u.PrimaryLanguage); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, err // an error occurred
	}

	return &u, nil // User found
}

func (conn *Conn) FindUserIdByEmail(ctx context.Context, email string) ([]byte, error) {
	var u models.User

	// Search for user by primary email
	row := conn.QueryRow(ctx, query.FindUserIdByPrimaryEmail, email)
	if err := row.Scan(&u.Id); err != nil {
		if err != sql.ErrNoRows {
			return nil, err // an error occurred
		}
		// If not found, search in UserEmails
		row = conn.QueryRow(ctx, query.FindUserIdByEmail, email)
		if err := row.Scan(&u.Id); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // No result found
			}
			return nil, err // an error occurred
		}
	}

	return u.Id, nil // User found
}

func (conn *Conn) FindUserIdByProvider(ctx context.Context, provider models.ProviderInfo) ([]byte, error) {
	var u models.User

	// Search in OAuthProviders
	row := conn.QueryRow(ctx, query.FindOAuthProviderByProviderNameAndProviderId, provider.Name, provider.Id)
	if err := row.Scan(&u.Id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found in both tables
		}
		return nil, err // an error occurred
	}

	return u.Id, nil // providerId found
}

func (conn *Conn) FindUserPicturesByUserId(ctx context.Context, userId []byte) ([]*models.UserPicture, error) {
	var userPictures []*models.UserPicture

	rows, err := conn.Query(ctx, query.FindUserPictureByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userPicture models.UserPicture
		if err := rows.Scan(&userPicture.Id, &userPicture.PictureType, &userPicture.PictureUrl, &userPicture.UserId); err != nil {
			return nil, err
		}
		userPictures = append(userPictures, &userPicture)
	}

	return userPictures, nil
}

func (conn *Conn) FindUserEmailsByUserId(ctx context.Context, userId []byte) ([]*models.UserEmail, error) {
	var userEmails []*models.UserEmail

	rows, err := conn.Query(ctx, query.FindUserEmailByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userEmail models.UserEmail
		if err := rows.Scan(&userEmail.Email, &userEmail.IsPrimary, &userEmail.Verified, &userEmail.UserId); err != nil {
			return nil, err
		}
		userEmails = append(userEmails, &userEmail)
	}

	return userEmails, nil
}

func (conn *Conn) FindUserOAuthProvidersByUserId(ctx context.Context, userId []byte) ([]*models.OAuthProvider, error) {
	var userProviders []*models.OAuthProvider

	rows, err := conn.Query(ctx, query.FindOAuthProviderByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userProvider models.OAuthProvider
		if err := rows.Scan(&userProvider.Id, &userProvider.ProviderName, &userProvider.ProviderId, &userProvider.Principal, &userProvider.Token, &userProvider.UserId); err != nil {
			return nil, err
		}
		userProviders = append(userProviders, &userProvider)
	}

	return userProviders, nil
}
