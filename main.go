package main

import (
	"log"

	"github.com/gophersnacks/gbfm/actions"
	"github.com/gophersnacks/gbfm/actions/content"
	"github.com/gophersnacks/gbfm/actions/snacks"
)

func main() {

	gbfmApp := actions.GBFMApp()
	snacksApp := snacks.App()
	contentApp, contentClose := content.App()

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
