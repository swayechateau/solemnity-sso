package database

import (
	"sso/internal/database/models"
	"sso/pkg/database/query"

	"github.com/jackc/pgx/v5"
	"github.com/swayedev/way"
)

// User
func GetUser(w *way.Context, id [16]byte) (*models.User, error) {
	u, err := FindUserById(w, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	// get user emails
	emails, err := FindUserEmailsByUserId(w, id)
	if err != nil {
		return nil, err
	}
	u.Emails = make([]models.UserEmail, len(emails))
	for i, email := range emails {
		u.Emails[i] = *email
	}

	// get user pictures
	pictures, err := FindUserPicturesByUserId(w, id)
	if err != nil {
		return nil, err
	}

	u.Pictures = make([]models.UserPicture, len(pictures))
	for i, picture := range pictures {
		u.Pictures[i] = *picture
	}

	// get user oauth providers
	providers, err := FindUserProvidersByUserId(w, id)
	if err != nil {
		return nil, err
	}

	u.Providers = make([]models.Provider, len(providers))
	for i, provider := range providers {
		u.Providers[i] = *provider
	}

	return u, nil
}

// Find User by Id
func FindUserById(w *way.Context, id [16]byte) (*models.User, error) {
	var u models.User
	ctx := w.Request.Context()

	row := w.PgxQueryRow(ctx, query.FindUserById, id)
	err := row.Scan(&u.Id, &u.Verified, &u.DisplayName, &u.PrimaryEmailHash, &u.PrimaryEmail, &u.PrimaryPictureId, &u.PrimaryLanguage, &u.CreatedAt, &u.UpdatedAt)
	if err == nil {
		return &u, nil // User found

	}
	if err == pgx.ErrNoRows {
		return nil, nil // No result found
	}
	return nil, err // an error occurred
}

func findUserIdInUserEmails(w *way.Context, email string) ([]byte, error) {
	ctx := w.Request.Context()
	var byteId []byte

	row := w.PgxQueryRow(ctx, query.FindUserIdByEmail, email)
	err := row.Scan(&byteId)
	if err != nil {
		return nil, err // an error occurred
	}

	return byteId, nil // User found
}

// Find UserId by Email
func FindUserIdByEmail(w *way.Context, email string) ([]byte, error) {
	var byteId []byte
	ctx := w.Request.Context()

	// Search for user by primary email
	row := w.PgxQueryRow(ctx, query.FindUserIdByPrimaryEmail, email)
	err := row.Scan(&byteId)
	if err == nil {
		return byteId, nil // User found
	}

	if err != pgx.ErrNoRows {
		return nil, err // an error occurred
	}

	// Search for user by email in UserEmails
	byteId, err = findUserIdInUserEmails(w, email)
	if err == nil {
		return byteId, nil // User found
	}

	if err != pgx.ErrNoRows {
		return nil, err // an error occurred
	}

	return nil, nil // User not found
}

// Find UserId by Provider
func FindUserIdByProvider(w *way.Context, name string, id string) ([]byte, error) {
	var byteId []byte
	ctx := w.Request.Context()

	// Search in OAuthProviders
	row := w.PgxQueryRow(ctx, query.FindUserIdByProviderNameAndId, name, id)
	err := row.Scan(&byteId)
	if err == nil {
		return byteId, nil // userId found
	}
	if err != pgx.ErrNoRows {
		return nil, err // an error occurred
	}
	return nil, nil // No result found
}

// Create User
func CreateUser(w *way.Context, u models.User) error {
	ctx := w.Request.Context()
	if err := u.Validate(); err != nil {
		return err
	}

	if err := w.PgxExecNoResult(
		ctx, query.CreateUser, u.Id,
		u.Verified, u.DisplayName,
		u.PrimaryEmailHash,
		u.PrimaryEmail, u.PrimaryPictureId,
		u.PrimaryLanguage,
	); err != nil {
		return err
	}

	email := models.UserEmail{
		UserId:    u.Id,
		EmailHash: u.PrimaryEmailHash,
		Email:     u.PrimaryEmail,
		Primary:   true,
		Verified:  u.Verified,
	}

	return CreateUserEmail(w, email)
}

// Update User
func UpdateUser(w *way.Context, u *models.User) error {
	ctx := w.Request.Context()
	// check if email exists in UserEmails

	// if not, create it

	// if yes, update it

	return w.PgxExecNoResult(ctx, query.UpdateUser, u.Verified, u.DisplayName, u.PrimaryEmailHash, u.PrimaryEmail, u.PrimaryPictureId, u.PrimaryLanguage, u.Id)
}

// Delete User
func DeleteUser(w *way.Context, id [16]byte) error {
	ctx := w.Request.Context()
	return w.PgxExecNoResult(ctx, query.DeleteUserById, id)
}
