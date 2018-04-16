package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gophersnacks/gbfm/pkg/render"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gophersnacks/gbfm/actions/gbfm"
	"github.com/gophersnacks/gbfm/models"
)

// GBFMApp is where all routes and middleware for gobeyond5minutes.com are defined
func GBFMApp() *buffalo.App {
	app := buffalo.New(buffalo.Options{
		Addr:        "0.0.0.0:3001",
		Env:         ENV,
		SessionName: "_gbfm_session",
	})
	// Automatically redirect to SSL
	app.Use(ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	}))

	if ENV == "development" {
		app.Use(middleware.ParameterLogger)
	}

	// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
	// Remove to disable this.
	app.Use(csrf.New)

	// Wraps each request in a transaction.
	//  c.Value("tx").(*pop.PopTransaction)
	// Remove to disable this.
	app.Use(middleware.PopTransaction(models.DB))

	// Setup and use translations:
	app.Use(translator.Middleware())
	gbfm.AddRoutes(app)

	admin := app.Group("/admin")
	admin.Use(SetCurrentUser)
	admin.Use(AdminAuthorize)
	admin.GET("/", AdminHandler)

	app.Use(SetCurrentUser)
	app.GET("/users/new", UsersNew)
	app.POST("/users", UsersCreate)
	app.GET("/signin", AuthNew)
	app.POST("/signin", AuthCreate)
	app.POST("/token", AuthToken)
	app.DELETE("/signout", AuthDestroy)
	app.GET("/authors", AuthorList)
	app.GET("/authors/{name}", AuthorShow)
	app.ServeFiles("/", render.AssetsBox) // serve files from the public directory

	return app
}
