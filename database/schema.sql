
CREATE TABLE IF NOT EXISTS Users (
    Id BINARY(16) PRIMARY KEY,
    Verified BOOLEAN,
    DisplayName VARCHAR(255),
    PrimaryEmail VARCHAR(255),
    PrimaryPictureId VARCHAR(255),
    PrimaryLanguage VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS UserPictures (
    Id VARCHAR(255) PRIMARY KEY,
    PictureType VARCHAR(50),
    PictureUrl VARCHAR(255),
    UserId BINARY(16)),
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS UserEmails (
    Email VARCHAR(255) PRIMARY KEY,
    IsPrimary BOOLEAN,
    Verified BOOLEAN,
    UserId BINARY(16),
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS OAuthProviders (
    Id VARCHAR(255) PRIMARY KEY,
    ProviderName VARCHAR(255),
    ProviderId VARCHAR(255),
    Principal VARCHAR(255),
    Token LONGTEXT,
    UserId BINARY(16),
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
    UNIQUE (ProviderName, ProviderId)
);

-- Clients
CREATE TABLE IF NOT EXISTS Clients (
    Id VARCHAR(255) PRIMARY KEY,
    ClientSecret VARCHAR(255),
    RedirectUri TEXT,
    Scopes TEXT,
    GrantTypes TEXT
);

-- Access Tokens
CREATE TABLE IF NOT EXISTS AccessTokens (
    TokenSignature VARCHAR(255) PRIMARY KEY,
    ClientId VARCHAR(255),
    TokenData BLOB,
    TokenExpiry TIMESTAMP,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE
);

-- Refresh Tokens
CREATE TABLE IF NOT EXISTS RefreshTokens (
    TokenSignature VARCHAR(255) PRIMARY KEY,
    ClientId VARCHAR(255),
    TokenData BLOB,
    TokenExpiry TIMESTAMP,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE
);

-- Authorization Codes
CREATE TABLE IF NOT EXISTS AuthCodes (
    CodeSignature VARCHAR(255) PRIMARY KEY,
    ClientId VARCHAR(255),
    CodeData BLOB,
    CodeExpiry TIMESTAMP,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE
);

-- User Consents
CREATE TABLE IF NOT EXISTS UserConsents (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    UserId BINARY(16),
    ClientId VARCHAR(255),
    Scopes TEXT,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
);

-- Scopes
CREATE TABLE IF NOT EXISTS Scopes (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    ScopeName VARCHAR(255) UNIQUE,
    description TEXT
);


-- INSERT INTO Users (id, user) VALUES (UNHEX(REPLACE(UUID(), '-', '')), 'Jimmy');
-- SELECT HEX(id) as id, user FROM Users WHERE HEX(id) = 'YOUR_UUID_HERE';
