package config

import "sso/pkg/env"

type AuthProviderConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUrl  string `json:"redirect_url"`
}

func GetGoogleConfig() AuthProviderConfig {
	return AuthProviderConfig{
		ClientId:     env.Check("GOOGLE_CLIENT_ID"),
		ClientSecret: env.Check("GOOGLE_CLIENT_SECRET"),
		RedirectUrl:  env.SetDefault("GOOGLE_REDIRECT_URL", "http://localhost:8080/oauth2/google/callback"),
	}
}

func GetGithubConfig() AuthProviderConfig {
	return AuthProviderConfig{
		ClientId:     env.Check("GITHUB_CLIENT_ID"),
		ClientSecret: env.Check("GITHUB_CLIENT_SECRET"),
		RedirectUrl:  env.SetDefault("GITHUB_REDIRECT_URL", "http://localhost:8080/oauth2/github/callback"),
	}
}

func GetMicrosoftConfig() AuthProviderConfig {
	return AuthProviderConfig{
		ClientId:     env.Check("MICROSOFT_CLIENT_ID"),
		ClientSecret: env.Check("MICROSOFT_CLIENT_SECRET"),
		RedirectUrl:  env.SetDefault("MICROSOFT_REDIRECT_URL", "http://localhost:8080/oauth2/microsoft/callback"),
	}
}

func GetUrl(route string) string {
	url := "http://localhost:" + GetPort() + "/"
	return env.SetDefault("URL", url)
}

func GetPort() string {
	return env.SetDefault("PORT", "8080")
}
