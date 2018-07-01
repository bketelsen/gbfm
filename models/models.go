package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gobuffalo/envy"
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
	pgHost := envy.Get("PG_HOST", "127.0.0.1")
	pgPortStr := envy.Get("PG_PORT", "5432")
	pgPort, err := strconv.Atoi(pgPortStr)
	if err != nil {
		log.Fatalf("invalid PG_PORT %s", pgPort)
	}
	pgUser := envy.Get("PG_USER", "postgres")
	pgSSL := envy.Get("PG_SSL", "disable")
	pgDB := envy.Get("PG_DB", "gbfm_development")
	pgPass := envy.Get("PG_PASS", "postgres")
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s sslmode=%s dbname=%s password=%s",
		pgHost,
		pgUser,
		pgSSL,
		pgDB,
		pgPass,
	)
	GORM, err = gorm.Open("postgres", connStr)
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
