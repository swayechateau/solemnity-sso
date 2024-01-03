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

// Create User Picture
func CreateUserPicture(w *way.Context, up models.UserPicture) error {
	return w.PgxExecNoResult(w.Request.Context(), query.CreateUserPicture, up.Id, up.Extension, up.Uri, up.UserId)
}

// Update User Picture
func UpdateUserPicture(w *way.Context, up *models.UserPicture) error {
	return w.PgxExecNoResult(w.Request.Context(), query.UpdateUserPicture, up.Extension, up.Uri, up.Id)
}

// Delete User Picture
func DeleteUserPicture(w *way.Context, id string) error {
	return w.PgxExecNoResult(w.Request.Context(), query.DeleteUserPictureById, id)
}
