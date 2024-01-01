package demo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sso/internal/database"
	"sso/internal/database/models"

	"github.com/google/uuid"
	"github.com/swayedev/way"
)

// Add Test User
func AddMe(w *way.Context) error {
	uuid := "7ccbce2f-3654-4497-8f62-7e11b89e98ce"
	pId := "990079e3-8327-4bf0-9ede-28e27ab22a9b"

	var nU models.User
	nU.Verified = true
	nU.SetDisplayName("Swaye Chateau")
	nU.SetPrimaryEmail("swaye@dev.com")
	nU.PrimaryLanguage = "en"
	nU.SetIdFromString(uuid)
	nU.SetPrimaryPictureIdFromString(pId)

	err := database.CreateUser(w, nU)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func AddUserHandler(c *way.Context) {

	y, _ := uuid.Parse("7ccbce2f-3654-4497-8f62-7e11b89e98ce")

	// Find user by id
	u, err := database.FindUserById(c, y)
	if err != nil {
		log.Printf("Error finding user: %v \n", err)
		return
	}

	if u == nil {
		err = AddMe(c)
		if err != nil {
			log.Printf("Error adding user: %v \n", err)
			return
		}

		c.Response.WriteHeader(http.StatusCreated)
		c.Response.Write([]byte("User not found, one was created"))
		return
	}

	c.Response.WriteHeader(http.StatusOK)
	// c.Response.Write([]byte("User found"))
	json.NewEncoder(c.Response).Encode(u.ToJson())
}
