-- Migration: Initial Schema
-- Created by: SolemnitySSO
-- Created on: 2023-12-26 04:56:00
-- Last Modified: 2023-12-26 04:56:00
-- Version: 0.9.0
-- Description: This migration deletes tables created by the initial schema for the SolemnitySSO OAuth Server.

-- Revert the addition of the Users table
DROP TABLE IF EXISTS Users;
-- Revert the addition of the User Pictures table
DROP TABLE IF EXISTS UserPictures;
-- Revert the addition of the User Emails table
DROP TABLE IF EXISTS UserEmails;

-- Revert the addition of the OAuth Providers table
DROP TABLE IF EXISTS Providers;

-- Revert the addition of the Clients table
DROP TABLE IF EXISTS Clients;
-- Revert the addition of the Access Tokens table
DROP TABLE IF EXISTS AccessTokens;
-- Revert the addition of the Refresh Tokens table
DROP TABLE IF EXISTS RefreshTokens;
-- Revert the addition of the Auth Codes table
DROP TABLE IF EXISTS AuthCodes;
-- Revert the addition of the User Consents table
DROP TABLE IF EXISTS UserConsents;
-- Revert the addition of the Scopes table
DROP TABLE IF EXISTS Scopes;
