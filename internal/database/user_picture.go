package database

import (
	"sso/internal/database/models"
	"sso/pkg/database/query"

	"github.com/swayedev/way"
)

// Finc User Pictures by UserId
func FindUserPicturesByUserId(w *way.Context, userId [16]byte) ([]*models.UserPicture, error) {
	var userPictures []*models.UserPicture
	ctx := w.Request.Context()

	rows, err := w.PgxQuery(ctx, query.FindUserPictureByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var up models.UserPicture
		if err := rows.Scan(&up.Id, &up.Extension, &up.Uri, &up.UserId, &up.CreatedAt, &up.UpdatedAt); err != nil {
			return nil, err
		}
		userPictures = append(userPictures, &up)
	}

	return userPictures, nil
}

// Find User Picture by Id
func FindUserPictureById(w *way.Context, id [16]byte) (*models.UserPicture, error) {
	var up models.UserPicture
	ctx := w.Request.Context()

	row := w.PgxQueryRow(ctx, query.FindUserPictureById, id)
	err := row.Scan(&up.Id, &up.Extension, &up.Uri, &up.UserId, &up.CreatedAt, &up.UpdatedAt)
	if err == nil {
		return &up, nil
	}

	return nil, err
}

// Create User Picture
func CreateUserPicture(w *way.Context, up models.UserPicture) error {
	if err := up.Validate(); err != nil {
		return err
	}
	return w.PgxExecNoResult(w.Request.Context(), query.CreateUserPicture, up.Id, up.Extension, up.Uri, up.UserId)
}

// Update User Picture
func UpdateUserPicture(w *way.Context, up *models.UserPicture) error {
	if err := up.Validate(); err != nil {
		return err
	}
	return w.PgxExecNoResult(w.Request.Context(), query.UpdateUserPicture, up.Extension, up.Uri, up.Id)
}

// Delete User Picture
func DeleteUserPicture(w *way.Context, id string) error {
	return w.PgxExecNoResult(w.Request.Context(), query.DeleteUserPictureById, id)
}
