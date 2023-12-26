package database

import (
	"context"
	"sso/database/models"
	"sso/database/query"
)

func (conn *Conn) UpdateUser(ctx context.Context, u *models.User) error {
	return conn.Insert(ctx, query.UpdateUser, u.Verified, u.DisplayName, u.PrimaryEmail, u.PrimaryPictureId, u.PrimaryLanguage, u.Id)
}

func (conn *Conn) UpdateUserPicture(ctx context.Context, up *models.UserPicture) error {
	return conn.Insert(ctx, query.UpdateUserPicture, up.PictureType, up.PictureUrl, up.Id)
}

func (conn *Conn) UpdateUserEmail(ctx context.Context, ue *models.UserEmail) error {
	return conn.Insert(ctx, query.UpdateUserEmail, ue.IsPrimary, ue.Verified, ue.Email)
}
