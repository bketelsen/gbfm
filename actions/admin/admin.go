package admin

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gophersnacks/gbfm/models"
	"github.com/qor/admin"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/media"
	"github.com/qor/session/manager"
)

// NewApp creates an admin app but does not start it
func NewApp(ctx context.Context) (*App, error) {
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

	models.GORM.AutoMigrate(&auth_identity.AuthIdentity{})

	log.Printf("migrating user model")
	models.GORM.AutoMigrate(&models.User{})
	media.RegisterCallbacks(models.GORM)

	// create GH auth impl. TODO: fill in the host
	ghAuth := newGHAuth(ghClientID, ghClientSecret, "")

	// Initalize
	assetFS, err := adminAssetFS()
	if err != nil {
		return nil, err
	}
	adm := admin.New(&admin.AdminConfig{
		DB:      models.GORM,
		Auth:    ghAuth,
		AssetFS: assetFS,
	})

	topic := adm.AddResource(&models.Topic{})
	topic.NewAttrs("-Slug")

	snack := adm.AddResource(&models.Snack{})
	snack.NewAttrs("-Slug")

	series := adm.AddResource(&models.Series{})
	series.NewAttrs("-Slug")

	episode := adm.AddResource(&models.Episode{})
	episode.Meta(&admin.Meta{Name: "Body", Type: "text"})

	guide := adm.AddResource(&models.Guide{})
	guide.NewAttrs("-Slug")

	u := adm.AddResource(&models.User{})
	u.NewAttrs("-Slug")

	author := adm.AddResource(&models.Author{})
	author.NewAttrs("-Slug")

	// mount the a mux to auth
	mux := http.NewServeMux()
	mux.Handle("/", ghAuth)
	// Admin.MountTo("/auth", mux)
	adm.MountTo("/admin", mux)

	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}

	return &App{
		ctx:     ctx,
		handler: manager.SessionManager.Middleware(mux),
	}, nil
}
