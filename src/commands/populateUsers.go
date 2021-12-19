package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	createAdmin()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}
}

func createAdmin() {
	ambassador := models.User{
		FirstName:    "vasa",
		LastName:     "petroff",
		Email:        "v@i.ua",
		IsAmbassador: false,
	}

	ambassador.SetPassword("password")

	database.DB.Create(&ambassador)
}
