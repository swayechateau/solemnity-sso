package database

import (
	"sso/internal/database/models"
	"sso/pkg/database/query"

	"github.com/swayedev/way"
)

// Find User Emails by UserId
func FindUserEmailsByUserId(w *way.Context, userId [16]byte) ([]*models.UserEmail, error) {
	var userEmails []*models.UserEmail
	ctx := w.Request.Context()

	rows, err := w.PgxQuery(ctx, query.FindUserEmailByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ue models.UserEmail
		if err := rows.Scan(&ue.EmailHash, &ue.Email, &ue.Primary, &ue.Verified, &ue.UserId, &ue.CreatedAt, &ue.UpdatedAt); err != nil {
			return nil, err
		}
		userEmails = append(userEmails, &ue)
	}

	return userEmails, nil
}

// Find User Email by Email
func FindUserEmailByEmail(w *way.Context, email string) (*models.UserEmail, error) {
	var ue models.UserEmail
	ctx := w.Request.Context()

	row := w.PgxQueryRow(ctx, query.FindUserEmailByEmail, email)
	err := row.Scan(&ue.EmailHash, &ue.Email, &ue.Primary, &ue.Verified, &ue.UserId, &ue.CreatedAt, &ue.UpdatedAt)
	if err == nil {
		return &ue, nil
	}

	return nil, err
}

// Create User Email
func CreateUserEmail(w *way.Context, ue models.UserEmail) error {
	if err := ue.Validate(); err != nil {
		return err
	}
	return w.PgxExecNoResult(w.Request.Context(), query.CreateUserEmail, ue.EmailHash, ue.Email, ue.Primary, ue.Verified, ue.UserId)
}

// Update User Email
func UpdateUserEmail(w *way.Context, ue *models.UserEmail) error {
	if err := ue.Validate(); err != nil {
		return err
	}
	return w.PgxExecNoResult(w.Request.Context(), query.UpdateUserEmail, ue.Primary, ue.Verified, ue.EmailHash)
}

func UpdateUserEmailIsPrimary(w *way.Context, email string, primary bool) error {
	return w.PgxExecNoResult(w.Request.Context(), query.UpdateUserEmailIsPrimary, primary, email)
}

// Delete User Email
func DeleteUserEmail(w *way.Context, email string) error {
	return w.PgxExecNoResult(w.Request.Context(), query.DeleteUserEmailByEmail, email)
}
