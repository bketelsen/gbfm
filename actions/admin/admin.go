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
	const (
		// Gophersnacks test 1
		ghClientID     = "3f53c147b90f7b5725db"
		ghClientSecret = "596000348d8467d2dfb42c5d0b26780d02156cf0"

		// Gophersnacks test 2
		// ghClientID     = "86deba66601689994258"
		// ghClientSecret = "4849e71a9026b6aa9232e29c3ee0929f214bf1ab"

		// Gophersnacks test 3
		// ghClientID = "9d2091f2ced29dd6a663"
		// ghClientSecret = "acec4b3eae80de53b494e746d3412a1cb5c183ae"
	)

	models.DB.AutoMigrate(&auth_identity.AuthIdentity{})

	models.DB.AutoMigrate(&models.User{})
	media.RegisterCallbacks(models.DB)

	// create GH auth inmpl. TODO: fill in the host
	provider := newGHProvider(ghClientID, ghClientSecret, "")

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
	mux.Handle("/", provider)
	// Admin.MountTo("/auth", mux)
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}
	http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
}
