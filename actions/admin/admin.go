package admin

import (
	"fmt"
	"net/http"

	"github.com/gophersnacks/gbfm/models"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/media"

	"github.com/qor/auth/auth_identity"
	"github.com/qor/session/manager"
)

var Auth = auth.New(&auth.Config{
	DB: models.DB,
	// User model needs to implement qor.CurrentUser interface (https://godoc.org/github.com/qor/qor#CurrentUser) to use it in QOR Admin
	UserModel: models.User{},
})

func Admin() {

	models.DB.AutoMigrate(&auth_identity.AuthIdentity{})

	models.DB.AutoMigrate(&models.User{})
	media.RegisterCallbacks(models.DB)

	// Register auth0 auth provider
	// TODO: GH key, etc...
	provider := newAuth0Provider("", "", "")
	Auth.RegisterProvider(provider)

	// Initalize
	Admin := admin.New(&admin.AdminConfig{
		DB:   models.DB,
		Auth: provider,
	})

	topic := Admin.AddResource(&models.Topic{})
	topic.NewAttrs("-Slug")

	snack := Admin.AddResource(&models.Snack{})
	snack.NewAttrs("-Slug")

	series := Admin.AddResource(&models.Series{})
	series.NewAttrs("-Slug")

	episode := Admin.AddResource(&models.Episode{})
	episode.Meta(&admin.Meta{Name: "Body", Type: "text"})

	guide := Admin.AddResource(&models.Guide{})
	guide.NewAttrs("-Slug")

	u := Admin.AddResource(&models.User{})
	u.NewAttrs("-Slug")

	author := Admin.AddResource(&models.Author{})
	author.NewAttrs("-Slug")

	// mount the a mux to auth
	mux := http.NewServeMux()
	mux.Handle("/auth/", Auth.NewServeMux())
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}
	http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
}
