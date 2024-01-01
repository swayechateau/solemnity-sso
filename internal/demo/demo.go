package demo

import (
	"fmt"
	"sso/internal/database"
	"sso/pkg/database/models"

	"github.com/swayedev/way"
)

// Add Test User
func AddMe(w way.Context) error {
	uuid := "7ccbce2f-3654-4497-8f62-7e11b89e98ce"
	pId := "990079e3-8327-4bf0-9ede-28e27ab22a9b"

	nU := models.User{
		Verified:        true,
		DisplayName:     "Swaye Chateau",
		PrimaryEmail:    "swaye@dev.com",
		PrimaryLanguage: "en",
	}

	nU.SetIdFromString(uuid)
	nU.SetPrimaryPictureIdFromString(pId)

	err := database.CreateUser(w, nU)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
