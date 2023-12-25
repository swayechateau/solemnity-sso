package config

import (
	"fmt"
	"os"
)

type AuthConfig struct {
	Google    AuthProviderConfig `json:"google"`
	Github    AuthProviderConfig `json:"github"`
	Microsoft AuthProviderConfig `json:"microsoft"`
}

type AuthProviderConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUrl  string `json:"redirect_url"`
}

func GetConfig() AuthConfig {
	return AuthConfig{
		Google:    GetGoogleConfig(),
		Github:    GetGithubConfig(),
		Microsoft: GetMicrosoftConfig(),
	}
}

func GetGoogleConfig() AuthProviderConfig {
	return AuthProviderConfig{
		ClientId:     checkEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: checkEnv("GOOGLE_CLIENT_SECRET"),
		RedirectUrl:  setUrl("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
	}
}

func GetGithubConfig() AuthProviderConfig {
	return AuthProviderConfig{
		ClientId:     checkEnv("GITHUB_CLIENT_ID"),
		ClientSecret: checkEnv("GITHUB_CLIENT_SECRET"),
		RedirectUrl:  setUrl("GITHUB_REDIRECT_URL", "http://localhost:8080/auth/github/callback"),
	}
}

func GetMicrosoftConfig() AuthProviderConfig {
	return AuthProviderConfig{
		ClientId:     checkEnv("MICROSOFT_CLIENT_ID"),
		ClientSecret: checkEnv("MICROSOFT_CLIENT_SECRET"),
		RedirectUrl:  setUrl("MICROSOFT_REDIRECT_URL", "http://localhost:8080/auth/microsoft/callback"),
	}
}

func checkEnv(env string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}

	tail := "not set... I hope you know what you're doing!"
	fmt.Printf(env + " " + tail)
	return ""
}

func setUrl(env string, url string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}

	tail := "not set... Using: " + url
	fmt.Printf(env + " " + tail)
	return url
}
