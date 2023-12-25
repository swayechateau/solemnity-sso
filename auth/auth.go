package auth

import (
	"github.com/gorilla/mux"
)

var randomStateString = "state" // make([]byte, 16)

func AuthHandler(r *mux.Router) *mux.Router {
	r = GoogleHandler(r)
	r = GithubHandler(r)
	r = MicrosoftHandler(r)
	r = QRCodeHandler(r)
	return r
}
