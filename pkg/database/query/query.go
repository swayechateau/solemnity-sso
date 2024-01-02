package query

var (
	// Selection queries
	FindUserIdByPrimaryEmail = "SELECT Id FROM Users WHERE PrimaryEmail = $1"
	FindUserIdByEmail        = "SELECT UserId FROM UserEmails WHERE Email = $1"
	FindUserById             = "SELECT * FROM Users WHERE Id = $1"
	FindUserPictureById      = "SELECT * FROM UserPictures WHERE Id = $1"
	// Gets array of UserPictures
	FindUserPictureByUserId = "SELECT * FROM UserPictures WHERE UserId = $1"
	FindUserEmailByEmail    = "SELECT * FROM UserEmails WHERE Email = $1"
	// Gets array of UserEmails
	FindUserEmailByUserId   = "SELECT * FROM UserEmails WHERE UserId = $1"
	FindProviderById        = "SELECT * FROM Providers WHERE Id = $1"
	FindProviderByNameAndId = "SELECT * FROM Providers WHERE ProviderName = $1 AND ProviderId = $2"
	// Gets array of Providers
	FindProviderByUserId       = "SELECT * FROM Providers WHERE UserId = $1"
	FindClientById             = "SELECT * FROM Clients WHERE Id = $1"
	FindAccessTokenBySignature = "SELECT * FROM AccessTokens WHERE TokenSignature = $1"
	// Gets array of AccessTokens
	FindAccessTokenByClientId   = "SELECT * FROM AccessTokens WHERE ClientId = $1"
	FindRefreshTokenBySignature = "SELECT * FROM RefreshTokens WHERE TokenSignature = $1"
	// Gets array of RefreshTokens
	FindRefreshTokenByClientId = "SELECT * FROM RefreshTokens WHERE ClientId = $1"
	FindAuthCodeBySignature    = "SELECT * FROM AuthCodes WHERE CodeSignature = $1"
	// Gets array of AuthCodes
	FindAuthCodeByClientId = "SELECT * FROM AuthCodes WHERE ClientId = $1"
	FindUserConsentById    = "SELECT * FROM UserConsents WHERE Id = $1"
	// Gets array of UserConsents
	FindUserConsentByUserId = "SELECT * FROM UserConsents WHERE UserId = $1"
	// Gets array of UserConsents
	FindUserConsentByClientId = "SELECT * FROM UserConsents WHERE ClientId = $1"
	FindScopeById             = "SELECT * FROM Scopes WHERE Id = $1"
	FindScopeByName           = "SELECT * FROM Scopes WHERE Name = $1"

	// Creation queries

	CreateUser         = "INSERT INTO Users (Id, Verified, DisplayName, PrimaryEmail, PrimaryPictureId, PrimaryLanguage) VALUES ($1, $2, $3, $4, $5, $6)"
	CreateUserPicture  = "INSERT INTO UserPictures (Id, PictureType, PictureUrl, UserId) VALUES ($1, $2, $3, $4)"
	CreateUserEmail    = "INSERT INTO UserEmails (Email, IsPrimary, Verified, UserId) VALUES ($1, $2, $3, $4)"
	CreateProvider     = "INSERT INTO Providers (ProviderName, ProviderId, Principal, UserId) VALUES ($1, $2, $3, $4)"
	CreateClient       = "INSERT INTO Clients (Id, ClientSecret, RedirectUris, Scopes, GrantTypes) VALUES ($1, $2, $3, $4, $5)"
	CreateAccessToken  = "INSERT INTO AccessTokens (TokenSignature, ClientId, TokenData, TokenExpiry) VALUES ($1, $2, $3, $4)"
	CreateRefreshToken = "INSERT INTO RefreshTokens (TokenSignature, ClientId, TokenData, TokenExpiry) VALUES ($1, $2, $3, $4)"
	CreateAuthCode     = "INSERT INTO AuthCodes (CodeSignature, ClientId, CodeData, CodeExpiry) VALUES ($1, $2, $3, $4)"
	CreateUserConsent  = "INSERT INTO UserConsents (UserId, ClientId, Scopes) VALUES ($1, $2, $3)"
	CreateScope        = "INSERT INTO Scopes (ScopeName, ScopeDescription) VALUES ($1, $2)"

	// Update queries
	UpdateUser        = "UPDATE Users SET Verified = $1, DisplayName = $2, PrimaryEmail = $3, PrimaryPictureId = $4, PrimaryLanguage = $5 WHERE Id = $6"
	UpdateUserPicture = "UPDATE UserPictures SET PictureType = $1, PictureUrl = $2 WHERE Id = $3"
	UpdateUserEmail   = "UPDATE UserEmails SET IsPrimary = $1, Verified = $2 WHERE Email = $3"

	// Deletion queries
	DeleteUserById          = "DELETE FROM Users WHERE Id = $1"
	DeleteUserPictureById   = "DELETE FROM UserPictures WHERE Id = $1"
	DeleteUserEmailByEmail  = "DELETE FROM UserEmails WHERE Email = $1"
	DeleteProviderById      = "DELETE FROM Providers WHERE Id = $1"
	DeleteClientById        = "DELETE FROM Clients WHERE Id = $1"
	DeleteAccessTokenBySig  = "DELETE FROM AccessTokens WHERE TokenSignature = $1"
	DeleteRefreshTokenBySig = "DELETE FROM RefreshTokens WHERE TokenSignature = $1"
	DeleteAuthCodeBySig     = "DELETE FROM AuthCodes WHERE CodeSignature = $1"
	DeleteUserConsentById   = "DELETE FROM UserConsents WHERE Id = $1"
	DeleteScopeById         = "DELETE FROM Scopes WHERE Id = $1"
)
