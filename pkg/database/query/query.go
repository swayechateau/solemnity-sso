package query

var (
	// Selection queries
	FindUserIdByPrimaryEmail = "SELECT Id FROM Users WHERE PrimaryEmail = ?"
	FindUserIdByEmail        = "SELECT UserId FROM UserEmails WHERE Email = ?"
	FindUserById             = "SELECT * FROM Users WHERE Id = ?"
	FindUserPictureById      = "SELECT * FROM UserPictures WHERE Id = ?"
	// Gets array of UserPictures
	FindUserPictureByUserId = "SELECT * FROM UserPictures WHERE UserId = ?"
	FindUserEmailByEmail    = "SELECT * FROM UserEmails WHERE Email = ?"
	// Gets array of UserEmails
	FindUserEmailByUserId   = "SELECT * FROM UserEmails WHERE UserId = ?"
	FindProviderById        = "SELECT * FROM Providers WHERE Id = ?"
	FindProviderByNameAndId = "SELECT * FROM Providers WHERE ProviderName = ? AND ProviderId = ?"
	// Gets array of Providers
	FindProviderByUserId       = "SELECT * FROM Providers WHERE UserId = ?"
	FindClientById             = "SELECT * FROM Clients WHERE Id = ?"
	FindAccessTokenBySignature = "SELECT * FROM AccessTokens WHERE TokenSignature = ?"
	// Gets array of AccessTokens
	FindAccessTokenByClientId   = "SELECT * FROM AccessTokens WHERE ClientId = ?"
	FindRefreshTokenBySignature = "SELECT * FROM RefreshTokens WHERE TokenSignature = ?"
	// Gets array of RefreshTokens
	FindRefreshTokenByClientId = "SELECT * FROM RefreshTokens WHERE ClientId = ?"
	FindAuthCodeBySignature    = "SELECT * FROM AuthCodes WHERE CodeSignature = ?"
	// Gets array of AuthCodes
	FindAuthCodeByClientId = "SELECT * FROM AuthCodes WHERE ClientId = ?"
	FindUserConsentById    = "SELECT * FROM UserConsents WHERE Id = ?"
	// Gets array of UserConsents
	FindUserConsentByUserId = "SELECT * FROM UserConsents WHERE UserId = ?"
	// Gets array of UserConsents
	FindUserConsentByClientId = "SELECT * FROM UserConsents WHERE ClientId = ?"
	FindScopeById             = "SELECT * FROM Scopes WHERE Id = ?"
	FindScopeByName           = "SELECT * FROM Scopes WHERE Name = ?"

	// Creation queries
	CreateUser         = "INSERT INTO Users (Id, Verified, DisplayName, PrimaryEmail, PrimaryPictureId, PrimaryLanguage) VALUES (?, ?, ?, ?, ?, ?)"
	CreateUserPicture  = "INSERT INTO UserPictures (Id, PictureType, PictureUrl, UserId) VALUES (?, ?, ?, ?)"
	CreateUserEmail    = "INSERT INTO UserEmails (Email, IsPrimary, Verified, UserId) VALUES (?, ?, ?, ?)"
	CreateProvider     = "INSERT INTO Providers (Id, ProviderName, ProviderId, Principal, Token, UserId) VALUES (?, ?, ?, ?, ?, ?)"
	CreateClient       = "INSERT INTO Clients (Id, ClientName, ClientSecret, RedirectUri, GrantTypes, ResponseTypes, Scope, OwnerId) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	CreateAccessToken  = "INSERT INTO AccessTokens (TokenSignature, ClientId, UserId, Scope, ExpiresIn, RequestedAt, GrantedAt, AccessData) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	CreateRefreshToken = "INSERT INTO RefreshTokens (TokenSignature, ClientId, UserId, Scope, ExpiresIn, RequestedAt, GrantedAt, AccessData) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	CreateAuthCode     = "INSERT INTO AuthCodes (CodeSignature, ClientId, UserId, Scope, ExpiresIn, RequestedAt, GrantedAt, AccessData) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	CreateUserConsent  = "INSERT INTO UserConsents (Id, UserId, ClientId, Scopes) VALUES (?, ?, ?, ?)"
	CreateScope        = "INSERT INTO Scopes (Id, ScopeName, ScopeDescription) VALUES (?, ?, ?)"

	// Update queries
	UpdateUser        = "UPDATE Users SET Verified = ?, DisplayName = ?, PrimaryEmail = ?, PrimaryPictureId = ?, PrimaryLanguage = ? WHERE Id = ?"
	UpdateUserPicture = "UPDATE UserPictures SET PictureType = ?, PictureUrl = ? WHERE Id = ?"
	UpdateUserEmail   = "UPDATE UserEmails SET IsPrimary = ?, Verified = ? WHERE Email = ?"

	// Deletion queries
	DeleteUserById          = "DELETE FROM Users WHERE Id = ?"
	DeleteUserPictureById   = "DELETE FROM UserPictures WHERE Id = ?"
	DeleteUserEmailByEmail  = "DELETE FROM UserEmails WHERE Email = ?"
	DeleteProviderById      = "DELETE FROM Providers WHERE Id = ?"
	DeleteClientById        = "DELETE FROM Clients WHERE Id = ?"
	DeleteAccessTokenBySig  = "DELETE FROM AccessTokens WHERE TokenSignature = ?"
	DeleteRefreshTokenBySig = "DELETE FROM RefreshTokens WHERE TokenSignature = ?"
	DeleteAuthCodeBySig     = "DELETE FROM AuthCodes WHERE CodeSignature = ?"
	DeleteUserConsentById   = "DELETE FROM UserConsents WHERE Id = ?"
	DeleteScopeById         = "DELETE FROM Scopes WHERE Id = ?"
)
