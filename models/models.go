package models

import (
	"log"

	"github.com/gobuffalo/pop"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a connection to your database to be used
// throughout your application.
//
// NOTE: This is not used. It's only here so that buffalo build can use it when it creates
// a packr box to run migrations...
var DB *pop.Connection
var GORM *gorm.DB

func init() {
	var err error
	//env := envy.Get("GO_ENV", "development")
	// TODO(BJK) - fix for environments
	GORM, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres sslmode=disable dbname=gbfm_development password=postgres")
	if err != nil {
		log.Fatal(err)
	}
	GORM.AutoMigrate(
		&Snack{},
		&Topic{},
		&Author{},
		&Episode{},
		&Guide{},
		&Series{},
		&User{},
		&AuthIdentity{},
		&Address{},
		&Author{},
	)
}

func sluggify(s string) string {
	return slug.Make(s)
}
