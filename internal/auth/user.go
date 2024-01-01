package auth

import (
	"encoding/json"
	"net/http"

	"github.com/swayedev/way"
)

func UserHandler(w *way.Context) {
	// Set the Content-Type header
	w.Response.Header().Set("Content-Type", "application/json")

	// Get the bearer token from the request header
	bearerToken := w.Request.Header.Get("Authorization")
	if bearerToken == "" {
		// If the token is missing, Unauthorized error
		http.Error(w.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := "user"
	// Write it back to the response
	json.NewEncoder(w.Response).Encode(user)
}
