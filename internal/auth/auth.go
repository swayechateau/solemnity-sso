package auth

import "github.com/swayedev/way"

func AuthHandler(w *way.Way) {
	// default oauth2 endpoints
	// w.GET("/oauth2/authorize", GoogleLoginHandler)
	// w.GET("/oauth2/token", GoogleCallbackHandler)
	// w.GET("/oauth2/revoke", GoogleLoginHandler)
	// w.GET("/oauth2/refreshtoken", GoogleCallbackHandler)
	// w.GET("/oauth2/user", GoogleCallbackHandler)
	//
	w.GET("/oauth2/google", GoogleLoginHandler)
	w.GET("/oauth2/google/callback", GoogleCallbackHandler)

	w.GET("/oauth2/github", GithubLoginHandler)
	w.GET("/oauth2/github/callback", GithubCallbackHandler)

	w.GET("/oauth2/microsoft", MicrosoftLoginHandler)
	w.GET("/oauth2/microsoft/callback", MicrosoftCallbackHandler)

	// w.GET("/oauth2/qr", QRCodeLoginHandler)
	// w.GET("/oauth2/qr/callback", QRCodeCallbackHandler)
}
