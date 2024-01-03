-- Migration: Initial Schema
-- Created by: SolemnitySSO
-- Created on: 2023-12-26 04:52:00
-- Last Modified: 2023-01-02 05:56:00
-- Version: 0.11.0
-- Description: This migration creates the initial schema for the SolemnitySSO OAuth Server.

-- Create the Users table
CREATE TABLE IF NOT EXISTS Users (
    Id UUID PRIMARY KEY,
    Verified BOOLEAN,
    DisplayName VARCHAR(255),
    PrimaryEmailHash VARCHAR(255) NOT NULL UNIQUE, -- This is the email address hash
    PrimaryEmailAddress VARCHAR(255) NOT NULL UNIQUE, -- This is the encrypted email address
    PrimaryPictureId UUID,
    PrimaryLanguage VARCHAR(255),
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the User Pictures table
CREATE TABLE IF NOT EXISTS UserPictures (
    Id UUID PRIMARY KEY,
    Extension VARCHAR(50),
    Uri VARCHAR(255) NOT NULL,
    UserId UUID,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the User Emails table
CREATE TABLE IF NOT EXISTS UserEmails (
    EmailHash VARCHAR(255) PRIMARY KEY, -- This is the email address hash
    EmailAddress VARCHAR(255) NOT NULL, -- This is the encrypted email address
    IsPrimary BOOLEAN,
    Verified BOOLEAN,
    UserId UUID,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the OAuth Providers table
CREATE TABLE IF NOT EXISTS Providers (
    Id SERIAL PRIMARY KEY,
    ProviderName VARCHAR(255),
    ProviderId VARCHAR(255),
    ProviderIdHash VARCHAR(255),
    Principal VARCHAR(255),
    UserId UUID,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
    UNIQUE (ProviderName, ProviderIdHash),
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the Clients table
CREATE TABLE IF NOT EXISTS Clients (
    Id VARCHAR(255) PRIMARY KEY,
    ClientSecret VARCHAR(255),
    RedirectUris TEXT,
    Scopes TEXT,
    GrantTypes TEXT,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the Access Tokens table
CREATE TABLE IF NOT EXISTS AccessTokens (
    TokenSignature VARCHAR(255) PRIMARY KEY,
    ClientId VARCHAR(255),
    TokenData BYTEA,
    TokenExpiry TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the Refresh Tokens table
CREATE TABLE IF NOT EXISTS RefreshTokens (
    TokenSignature VARCHAR(255) PRIMARY KEY,
    ClientId VARCHAR(255),
    TokenData BYTEA,
    TokenExpiry TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the Auth Codes table
CREATE TABLE IF NOT EXISTS AuthCodes (
    CodeSignature VARCHAR(255) PRIMARY KEY,
    ClientId VARCHAR(255),
    CodeData BYTEA,
    CodeExpiry TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the User Consents table
CREATE TABLE IF NOT EXISTS UserConsents (
    Id SERIAL PRIMARY KEY,
    UserId UUID,
    ClientId VARCHAR(255),
    Scopes TEXT,
    FOREIGN KEY (ClientId) REFERENCES Clients(Id) ON DELETE CASCADE,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the Scopes table
CREATE TABLE IF NOT EXISTS Scopes (
    Id SERIAL PRIMARY KEY,
    ScopeName VARCHAR(255) UNIQUE,
    ScopeDescription TEXT,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
