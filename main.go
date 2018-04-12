package main

import (
	"log"

	"github.com/gophersnacks/gbfm/actions"
)

func main() {

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
