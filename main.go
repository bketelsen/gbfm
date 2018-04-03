package main

import (
	"log"

	"github.com/gophersnacks/gbfm/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
