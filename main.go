package main

import (
	"log"
	"os"

	"github.com/gophersnacks/gbfm/actions"
	"github.com/gophersnacks/gbfm/models"
	"github.com/markbates/going/defaults"
)

func main() {
	baseURL := defaults.String(os.Getenv("CMS_URL"), "https://content.gophersnacks.com")
	models.BaseURL = baseURL

	gbfmApp := actions.GBFMApp()
	snacksApp := actions.SnacksApp()
	contentApp, contentClose := actions.ContentApp()

	errCh := make(chan error)
	go func() {
		if err := gbfmApp.Serve(); err != nil {
			log.Printf("ERROR: gbfm app crashed")
			errCh <- err
		}
	}()
	go func() {
		if err := snacksApp.Serve(); err != nil {
			log.Printf("ERROR: snacks app crashed")
			errCh <- err
		}
	}()
	go func() {
		defer contentClose()
		if err := contentApp.Serve(); err != nil {
			log.Printf("ERROR: content app crashed")
			errCh <- err
		}
	}()
	err := <-errCh
	log.Fatal(err)
}
