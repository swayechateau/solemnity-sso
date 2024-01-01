package auth

import (
	"encoding/json"
	"log"
	"sso/internal/database"

	"github.com/google/uuid"
	"github.com/swayedev/way"
)

func UserHandler(w *way.Context) {
	// Set the Content-Type header
	w.Response.Header().Set("Content-Type", "application/json")
	w.Response.Write([]byte("User Handler"))

	// Get the bearer token from the request header
	// bearerToken := w.Request.Header.Get("Authorization")
	// if bearerToken == "" {
	// 	// If the token is missing, Unauthorized error
	// 	http.Error(w.Response, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	// get user from db
	y, err := uuid.Parse("7ccbce2f-3654-4497-8f62-7e11b89e98ce")
	if err != nil {
		log.Printf("Error Passing User Id: %v \n", err)
	}
	u, err := database.FindUserById(w, y)
	if err != nil {
		log.Printf("Error Finding User: %v \n", err)
	}

	if u == nil {
		log.Printf("User Not Found")
		json.NewEncoder(w.Response).Encode("User Not Found")
		return
	}

	json.NewEncoder(w.Response).Encode(u.ToJson())
}
