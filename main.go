package main

import (
	"context"
	"log"

	"github.com/gophersnacks/gbfm/actions/admin"
	"github.com/gophersnacks/gbfm/actions/gbfm"
	"github.com/gophersnacks/gbfm/actions/snacks"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	snacksApp := snacks.App()
	snacksApp.Context = ctx
	gbfmApp := gbfm.App()
	gbfmApp.Context = ctx
	//	contentApp := content.App()

	go func() {
		if err := gbfmApp.Serve(); err != nil {
			log.Printf("ERROR: gbfm app crashed (%s)", err)
			cancel()
		}
	}()
	go func() {
		if err := snacksApp.Serve(); err != nil {
			log.Printf("ERROR: snacks app crashed (%s)", err)
			cancel()
		}
	}()
	go func() {
		admin.Admin()
		log.Printf("ERROR: admin app crashed (%s)")
		cancel()
	}()

	select {
	case <-ctx.Done():
		log.Fatal("shutting down")
	}
}
