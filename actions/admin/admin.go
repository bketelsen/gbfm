package admin

import (
	"fmt"
	"net/http"

	"github.com/gophersnacks/gbfm/models"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/media"

	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	"github.com/qor/qor"
	"github.com/qor/session/manager"
)

var Auth = auth.New(&auth.Config{
	DB: models.DB,
	// User model needs to implement qor.CurrentUser interface (https://godoc.org/github.com/qor/qor#CurrentUser) to use it in QOR Admin
	UserModel: models.User{},
})

type AdminAuth struct{}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
		}
		return qorCurrentUser
	}
	return nil
}

func Admin() {

	models.DB.AutoMigrate(&auth_identity.AuthIdentity{})

	models.DB.AutoMigrate(&models.User{})
	media.RegisterCallbacks(models.DB)
	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{}))

	// Initalize
	Admin := admin.New(&admin.AdminConfig{
		DB:   models.DB,
		Auth: &AdminAuth{},
	})

	// Allow to use Admin to manage User, Product

	Admin.AddResource(&models.Topic{})
	Admin.AddResource(&models.Snack{})
	Admin.AddResource(&models.Series{})
	episode := Admin.AddResource(&models.Episode{})
	episode.Meta(&admin.Meta{Name: "Body", Type: "text"})
	Admin.AddResource(&models.Guide{})
	Admin.AddResource(&models.User{})
	Admin.AddResource(&models.Author{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()
	mux.Handle("/auth/", Auth.NewServeMux())

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}
	http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
}
