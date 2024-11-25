package models

import (
	"asrs/app/models/auth"
	"github.com/RikiIhsan/lib/database"
)

// migrate a models to your database
func Migrate() {
	if err := database.Session["sdp"].DB.AutoMigrate(
		&auth.User{},
		&auth.Key{},
	); err != nil {
		panic(err)
	}
}
