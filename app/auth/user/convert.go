package user

import "sso/database/models"

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
