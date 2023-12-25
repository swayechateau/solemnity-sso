package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func QRCodeLoginHandler(w http.ResponseWriter, r *http.Request) { /* ... */ }

func QRCodeHandler(r *mux.Router) *mux.Router {
	r.HandleFunc("/auth/qr", QRCodeLoginHandler)
	return r
}
