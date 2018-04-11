package actions

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/gophersnacks/gbfm/content"
	"github.com/gophersnacks/gbfm/pkg/system/admin"
	"github.com/gophersnacks/gbfm/pkg/system/api/analytics"
	"github.com/gophersnacks/gbfm/pkg/system/db"
)

func ContentApp() error {

	db.Init()
	defer db.Close()

	analytics.Init()
	defer analytics.Close()

	admin.Run()

	// init search index
	go db.InitSearchIndex()

	// save the https port the system is listening on so internal system can make
	// HTTP api calls while in dev or production w/o adding more cli flags
	port := 8080
	err := db.PutConfig("http_port", fmt.Sprintf("%d", port))
	if err != nil {
		log.Fatalln("System failed to save config. Please try to run again.", err)
	}

	bind := "localhost"
	err = db.PutConfig("bind_addr", bind)
	if err != nil {
		log.Fatalln("System failed to save config. Please try to run again.", err)
	}

	fmt.Printf("Server listening at %s:%d for HTTP requests...\n", bind, port)
	fmt.Println("\nVisit '/admin' to get started.")
	return http.ListenAndServe(fmt.Sprintf("%s:%d", bind, port), nil)
}
