package database

import (
	"database/sql"
	"sso/internal/database/models"
	"sso/pkg/database/query"

	"github.com/swayedev/way"
)

// User

// Find User by Id
func FindUserById(w *way.Context, id [16]byte) (*models.User, error) {
	var u models.User
	ctx := w.Request.Context()
	var (
		byteId        []byte
		bytePictureId []byte
	)

	row := w.SqlQueryRow(ctx, query.FindUserById, id[:])
	if err := row.Scan(&byteId, &u.Verified, &u.DisplayName, &u.PrimaryEmail, &bytePictureId, &u.PrimaryLanguage, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, err // an error occurred
	}
	u.SetIdFromBytes(byteId)
	u.SetPrimaryPictureIdFromBytes(bytePictureId)
	return &u, nil // User found
}

func findUserIdInUserEmails(w *way.Context, email string) ([]byte, error) {
	ctx := w.Request.Context()
	var byteId []byte

	row := w.SqlQueryRow(ctx, query.FindUserIdByEmail, email)
	err := row.Scan(&byteId)
	if err != nil {
		return nil, err // an error occurred
	}

	return byteId, nil // User found
}

// Find UserId by Email
func FindUserIdByEmail(w *way.Context, email string) ([16]byte, error) {
	var (
		u      models.User
		byteId []byte
	)
	ctx := w.Request.Context()

	// Search for user by primary email
	row := w.SqlQueryRow(ctx, query.FindUserIdByPrimaryEmail, email)
	err := row.Scan(&byteId)
	u.SetIdFromBytes(byteId)
	if err == nil {
		return u.Id, nil // User found
	}

	if err != sql.ErrNoRows {
		return u.Id, err // an error occurred
	}

	// Search for user by email in UserEmails
	byteId, err = findUserIdInUserEmails(w, email)
	u.SetIdFromBytes(byteId)
	if err == nil {
		return u.Id, nil // User found
	}

	if err != sql.ErrNoRows {
		return u.Id, err // an error occurred
	}

	return u.Id, nil // User not found
}

// Find UserId by Provider
func FindUserIdByProvider(w *way.Context, name string, id string) ([16]byte, error) {
	var (
		u      models.User
		byteId []byte
	)
	ctx := w.Request.Context()

	// Search in OAuthProviders
	row := w.SqlQueryRow(ctx, query.FindProviderByNameAndId, name, id)
	if err := row.Scan(&byteId); err != nil {
		if err == sql.ErrNoRows {
			return u.Id, nil // No result found
		}
		return u.Id, err // an error occurred
	}
	u.SetIdFromBytes(byteId)
	return u.Id, nil // providerId found
}

// Create User
func CreateUser(w *way.Context, u models.User) error {
	ctx := w.Request.Context()
	if err := w.SqlExecNoResult(
		ctx, query.CreateUser, u.Id[:],
		u.Verified, u.DisplayName,
		u.PrimaryEmail, u.PrimaryPictureId[:],
		u.PrimaryLanguage,
	); err != nil {
		return err
	}

	email := models.UserEmail{
		UserId:   u.Id,
		Email:    u.PrimaryEmail,
		Primary:  true,
		Verified: u.Verified,
	}
	return CreateUserEmail(w, email)
}

// Update User
func UpdateUser(w *way.Context, u *models.User) error {
	ctx := w.Request.Context()
	// check if email exists in UserEmails

	// if not, create it

	// if yes, update it

	return w.SqlExecNoResult(ctx, query.UpdateUser, u.Verified, u.DisplayName, u.PrimaryEmail, u.PrimaryPictureId, u.PrimaryLanguage, u.Id)
}

// Delete User
func DeleteUser(w *way.Context, id [16]byte) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteUserById, id)
}

// User Emails

// Find User Emails by UserId
func FindUserEmailsByUserId(w *way.Context, userId [16]byte) ([]*models.UserEmail, error) {
	var userEmails []*models.UserEmail
	ctx := w.Request.Context()

	rows, err := w.SqlQuery(ctx, query.FindUserEmailByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userEmail models.UserEmail
		if err := rows.Scan(&userEmail.Email, &userEmail.Primary, &userEmail.Verified, &userEmail.UserId); err != nil {
			return nil, err
		}
		userEmails = append(userEmails, &userEmail)
	}

	return userEmails, nil
}

// Create User Email
func CreateUserEmail(w *way.Context, ue models.UserEmail) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateUserEmail, ue.Email, ue.Primary, ue.Verified, ue.UserId[:])
}

// Update User Email
func UpdateUserEmail(w *way.Context, ue *models.UserEmail) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.UpdateUserEmail, ue.Primary, ue.Verified, ue.Email)
}

// Delete User Email
func DeleteUserEmail(w *way.Context, email string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteUserEmailByEmail, email)
}

// User Pictures

// Finc User Pictures by UserId
func FindUserPicturesByUserId(w *way.Context, userId [16]byte) ([]*models.UserPicture, error) {
	var userPictures []*models.UserPicture
	ctx := w.Request.Context()

	rows, err := w.SqlQuery(ctx, query.FindUserPictureByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userPicture models.UserPicture
		if err := rows.Scan(&userPicture.Id, &userPicture.Type, &userPicture.Url, &userPicture.UserId); err != nil {
			return nil, err
		}
		userPictures = append(userPictures, &userPicture)
	}

	return userPictures, nil
}

// Create User Picture
func CreateUserPicture(w *way.Context, up models.UserPicture) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.CreateUserPicture, up.Id, up.Type, up.Url, up.UserId)
}

// Update User Picture
func UpdateUserPicture(w *way.Context, up *models.UserPicture) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.UpdateUserPicture, up.Type, up.Url, up.Id)
}

// Delete User Picture
func DeleteUserPicture(w *way.Context, id string) error {
	ctx := w.Request.Context()
	return w.SqlExecNoResult(ctx, query.DeleteUserPictureById, id)
}

// User OAuth Providers

// Find Providers by UserId
func FindUserProvidersByUserId(w *way.Context, userId [16]byte) ([]*models.Provider, error) {
	var userProviders []*models.Provider
	ctx := w.Request.Context()

	rows, err := w.SqlQuery(ctx, query.FindProviderByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userProvider models.Provider
		if err := rows.Scan(&userProvider.Id, &userProvider.Name, &userProvider.ProviderUserId, &userProvider.Principal, &userProvider.UserId); err != nil {
			return nil, err
		}
		userProviders = append(userProviders, &userProvider)
	}

	return userProviders, nil
}
