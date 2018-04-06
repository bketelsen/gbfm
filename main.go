package main

import (
	"log"
	"os"

	"github.com/gophersnacks/gbfm/actions"
	"github.com/gophersnacks/gbfm/models"
	"github.com/markbates/going/defaults"
)

func main() {
	app := actions.App()
	baseURL := defaults.String(os.Getenv("CMS_URL"), "https://content.gophersnacks.com")
	models.BaseURL = baseURL
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
