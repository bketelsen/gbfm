package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/gophersnacks/gbfm/actions/renderengine"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/gophersnacks/gbfm/actions/gbfm"
	"github.com/gophersnacks/gbfm/actions/snacks"
	"github.com/gophersnacks/gbfm/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
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
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())
		admin := app.Group("/admin")
		admin.Use(SetCurrentUser)
		admin.Use(AdminAuthorize)
		admin.GET("/", AdminHandler)

		// have caddy mux hostnames to these paths
		snacksGroup := app.Group("/snacks")
		snacks.AddRoutes(snacksGroup)

		gbfmGroup := app.Group("/gbfm")
		gbfm.AddRoutes(gbfmGroup)

		// gifm := app.Group("/gifm")

		app.Use(SetCurrentUser)
		app.GET("/", homeHandler)
		app.GET("/users/new", UsersNew)
		app.POST("/users", UsersCreate)
		app.GET("/signin", AuthNew)
		app.POST("/signin", AuthCreate)
		app.POST("/token", AuthToken)
		app.DELETE("/signout", AuthDestroy)
		app.GET("/authors", AuthorList)
		app.GET("/authors/{name}", AuthorShow)
		app.ServeFiles("/", renderengine.AssetsBox) // serve files from the public directory
	}

	return app
}
