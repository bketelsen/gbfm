package models

import (
	"log"

	"github.com/gophersnacks/gbfm/models/users"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *gorm.DB //*pop.Connection

func init() {
	var err error
	//env := envy.Get("GO_ENV", "development")
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=gbfm_development password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(
		&Snack{},
		&Topic{},
		&Author{},
		&Episode{},
		&Guide{},
		&Series{},
		&users.User{},
		&users.AuthIdentity{},
		&users.Address{},
	)
}
