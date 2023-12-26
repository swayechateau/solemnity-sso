package database

import (
	"context"
	"sso/database/models"
	"sso/database/query"
)

func (conn *Conn) CreateUser(ctx context.Context, u *models.User) error {
	err := conn.Insert(ctx, query.CreateUser, u.Id, u.Verified, u.DisplayName, u.PrimaryEmail, u.PrimaryPictureId, u.PrimaryLanguage)
	if err != nil {
		return err
	}

	email := models.UserEmail{
		UserId:    u.Id,
		Email:     u.PrimaryEmail,
		IsPrimary: true,
		Verified:  u.Verified,
	}
	return conn.CreateUserEmail(ctx, &email)
}

func (conn *Conn) CreateUserPicture(ctx context.Context, up *models.UserPicture) error {
	return conn.Insert(ctx, query.CreateUserPicture, up.Id, up.PictureType, up.PictureUrl, up.UserId)
}

func (conn *Conn) CreateUserEmail(ctx context.Context, ue *models.UserEmail) error {
	return conn.Insert(ctx, query.CreateUserEmail, ue.Email, ue.IsPrimary, ue.Verified, ue.UserId)
}

func (conn *Conn) CreateOAuthProvider(ctx context.Context, op *models.OAuthProvider) error {
	return conn.Insert(ctx, query.CreateOAuthProvider, op.Id, op.ProviderName, op.ProviderId, op.Principal, op.Token, op.UserId)
}

func (conn *Conn) CreateClient(ctx context.Context, c *models.Client) error {
	return conn.Insert(ctx, query.CreateClient, c)
}

func (conn *Conn) CreateAccessToken(ctx context.Context, at *models.AccessToken) error {
	return conn.Insert(ctx, query.CreateAccessToken, at)
}

func (conn *Conn) CreateRefreshToken(ctx context.Context, rt *models.RefreshToken) error {
	return conn.Insert(ctx, query.CreateRefreshToken, rt)
}

func (conn *Conn) CreateAuthCode(ctx context.Context, ac *models.AuthCode) error {
	return conn.Insert(ctx, query.CreateAuthCode, ac)
}

func (conn *Conn) CreateUserConsent(ctx context.Context, uc *models.UserConsent) error {
	return conn.Insert(ctx, query.CreateUserConsent, uc)
}

func (conn *Conn) CreateScope(ctx context.Context, s *models.Scope) error {
	return conn.Insert(ctx, query.CreateScope, s)
}
