package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gophersnacks/gbfm/actions/content"
	"github.com/gophersnacks/gbfm/actions/renderengine"
	"github.com/gophersnacks/gbfm/models"
	"github.com/unrolled/secure"
)

// ContentApp is where all routes and middleware for the admin interface are defined.
//
// The second parameter returned should be called by the caller in a defer
func ContentApp() (*buffalo.App, func()) {
	app := buffalo.New(buffalo.Options{
		Addr:        "0.0.0.0:8080",
		Env:         ENV,
		SessionName: "_admin_session",
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
	content.AddRoutes(app)
	app.ServeFiles("/", renderengine.AssetsBox) // serve files from the public directory

	return app, func() {} // TODO: remove the close func
}
