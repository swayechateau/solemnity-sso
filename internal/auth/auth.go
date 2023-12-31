package auth

import (
	"math/rand"

	"github.com/swayedev/way"
)

func AuthHandler(w *way.Way) {
	// default oauth2 endpoints

	// External OAuth2 providers
	w.GET("/oauth2/google", GoogleLoginHandler)
	w.GET("/oauth2/google/callback", GoogleCallbackHandler)

	w.GET("/oauth2/github", GithubLoginHandler)
	w.GET("/oauth2/github/callback", GithubCallbackHandler)

	w.GET("/oauth2/microsoft", MicrosoftLoginHandler)
	w.GET("/oauth2/microsoft/callback", MicrosoftCallbackHandler)
}

func randomCode() string {
	return randomString(randomBetween5And20())
}

func randomBetween5And20() int {
	return rand.Intn(16) + 5 // rand.Intn(16) gives a number between 0 and 15, so add 5
}

func randomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
