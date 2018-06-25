package main

import (
	"log"

	"github.com/gophersnacks/gbfm/actions/admin"
	"github.com/gophersnacks/gbfm/actions/gbfm"
	"github.com/gophersnacks/gbfm/actions/snacks"
)

func main() {

	snacksApp := snacks.App()
	gbfmApp := gbfm.App()
	//	contentApp := content.App()

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
		if err := admin.App(); err != nil {
			log.Printf("ERROR: admin app crashed")
			errCh <- err
		}
	}()
	err := <-errCh
	log.Fatal(err)
}
