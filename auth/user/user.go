package user

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sso/database"
	"sso/database/models"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type AuthUser struct {
	Id              uuid.UUID       `json:"id"`
	Verified        bool            `json:"verified"`
	DisplayName     string          `json:"display_name"`
	PrimaryEmail    string          `json:"primary_email"`
	PrimaryPicture  string          `json:"primary_picture"`
	PrimaryLanguage string          `json:"primary_language"`
	Pictures        []UserPicture   `json:"profile_pictures"`
	Email           []AuthUserEmail `json:"email"`
	OAuthProviders  []OAuthProvider `json:"oauth_providers"`
}

type UserPicture struct {
	Id   string `json:"id"`
	Type string `json:"type"` // gif, jpeg, png
	Url  string `json:"url"`
}

type AuthUserEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

type OAuthProvider struct {
	Id        string        `json:"id"`
	Provider  string        `json:"provider"`
	Principal string        `json:"principal"`
	Token     *oauth2.Token `json:"token"`
}

func (u *AuthUser) IdToByte() []byte {
	return u.Id[:]
}

func (u *AuthUser) IdToString() string {
	return u.Id.String()
}

func (u *AuthUser) ToUser() *models.User {
	return &models.User{
		Id:               u.IdToByte(),
		Verified:         u.Verified,
		DisplayName:      u.DisplayName,
		PrimaryEmail:     u.PrimaryEmail,
		PrimaryPictureId: u.PrimaryPicture,
		PrimaryLanguage:  u.PrimaryLanguage,
	}
}

func GetOAuthInfo(token *oauth2.Token, api string) ([]byte, error) {
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err.Error())
	}

	req.Header.Add("Authorization", "token "+token.AccessToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func GetUserByProvider(conn *database.Conn, ctx context.Context, provider *models.ProviderInfo) (*AuthUser, error) {
	id, err := conn.FindUserIdByProvider(ctx, *provider)
	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, nil
	}

	return GetUserById(conn, ctx, id)
}

func GetUserByEmail(conn *database.Conn, ctx context.Context, email string) (*AuthUser, error) {
	id, err := conn.FindUserIdByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, nil
	}

	return GetUserById(conn, ctx, id)
}

func GetUserById(conn *database.Conn, ctx context.Context, id []byte) (*AuthUser, error) {
	user, err := conn.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return ConvertUser(user), nil
}

func CreateUser(conn *database.Conn, ctx context.Context, user *models.User) error {
	return conn.CreateUser(ctx, user)
}

func AddOAuthProvider(conn *database.Conn, ctx context.Context, provider *models.OAuthProvider) error {
	return conn.CreateOAuthProvider(ctx, provider)
}

func ConvertUser(user *models.User) *AuthUser {
	return &AuthUser{
		Id:              user.IdToUUID(),
		Verified:        user.Verified,
		DisplayName:     user.DisplayName,
		PrimaryEmail:    user.PrimaryEmail,
		PrimaryPicture:  user.PrimaryPictureId,
		PrimaryLanguage: user.PrimaryLanguage,
	}
}

func GetAuthUser(conn *database.Conn, ctx context.Context, id []byte) (*AuthUser, error) {
	// get user
	user, err := conn.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	// get user pictures
	pictures, err := conn.FindUserPicturesByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	// get user emails
	emails, err := conn.FindUserEmailsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	// get user oauth providers
	providers, err := conn.FindUserOAuthProvidersByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	// convert user
	authUser := ConvertUser(user)

	// convert pictures
	authUser.Pictures = ConvertUserPictures(pictures)

	// convert emails
	authUser.Email = ConvertUserEmails(emails)

	// convert oauth providers
	authUser.OAuthProviders = ConvertOAuthProviders(providers)

	return authUser, nil

}

func ConvertUserPictures(uPs []*models.UserPicture) []UserPicture {
	pictures := make([]UserPicture, len(uPs))
	for i, picture := range uPs {
		pictures[i] = UserPicture{
			Id:   picture.Id,
			Type: picture.PictureType,
			Url:  picture.PictureUrl,
		}
	}
	return pictures
}

func ConvertUserEmails(uEs []*models.UserEmail) []AuthUserEmail {
	emails := make([]AuthUserEmail, len(uEs))
	for i, email := range uEs {
		emails[i] = AuthUserEmail{
			Email:    email.Email,
			Primary:  email.IsPrimary,
			Verified: email.Verified,
		}
	}
	return emails
}

func ConvertOAuthProviders(oPs []*models.OAuthProvider) []OAuthProvider {
	providers := make([]OAuthProvider, len(oPs))
	for i, provider := range oPs {
		providers[i] = OAuthProvider{
			Id:        provider.Id,
			Provider:  provider.ProviderName,
			Principal: provider.Principal,
		}
	}
	return providers
}
