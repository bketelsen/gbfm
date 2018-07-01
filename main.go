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

	snacksApp := snacks.App(ctx)
	gbfmApp := gbfm.App(ctx)
	adminApp, err := admin.NewApp(ctx)
	if err != nil {
		log.Fatalf("error creating new admin server (%s)", err)
	}

	errCh := make(chan error)
	go func() {
		log.Printf("starting gbfm app")
		if err := gbfmApp.Serve(); err != nil {
			log.Printf("ERROR: gbfm app crashed")
			errCh <- err
		}
	}()
	go func() {
		log.Printf("starting snacks app")
		if err := snacksApp.Serve(); err != nil {
			log.Printf("ERROR: snacks app crashed")
			errCh <- err
		}
	}()
	go func() {
		log.Printf("starting admin app")
		if err := adminApp.Serve(); err != nil {
			log.Printf("ERROR: admin app crashed")
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		cancel()
		log.Fatal(err)
	case <-ctx.Done():
		cancel()
		log.Fatalf("shit stopped (%s)", ctx.Err())
	}
}
