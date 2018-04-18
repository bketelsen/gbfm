package gbfm

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gophersnacks/gbfm/pkg/render"
	"github.com/gophersnacks/gbfm/pkg/web"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gophersnacks/gbfm/models"
)

var r = render.New("common_layout.html")

// App is where all routes and middleware for gobeyond5minutes.com are defined
func App() *buffalo.App {
	app := buffalo.New(buffalo.Options{
		Addr:        "0.0.0.0:3001",
		Env:         web.ENV,
		SessionName: "_gbfm_session",
	})
	// Automatically redirect to SSL
	app.Use(ssl.ForceSSL(secure.Options{
		SSLRedirect:     web.ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	}))

	if web.ENV == "development" {
		app.Use(middleware.ParameterLogger)
	}

	// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
	// Remove to disable this.
	app.Use(csrf.New)

	app.Use(web.LayoutMiddleware(
		"Go Beyond Five Minutes",
		"gbfm/partials/nav.html",
		"gbfm/partials/footer.html",
	))

	// Wraps each request in a transaction.
	//  c.Value("tx").(*pop.PopTransaction)
	// Remove to disable this.
	app.Use(middleware.PopTransaction(models.DB))

	// Setup and use translations:
	app.Use(web.Translator.Middleware())
	app.GET("/", homeHandler)

	app.GET("/authors", AuthorList)
	app.GET("/authors/{name}", AuthorShow)
	app.GET("/series", SeriesList)
	app.GET("/series/{name}", SeriesShow)
	app.GET("/guides", GuideList)
	app.GET("/guides/{name}", GuideShow)
	app.GET("/episodes", EpisodeList)
	app.GET("/episodes/{name}", EpisodeShow)
	app.ServeFiles("/", render.AssetsBox) // serve files from the public directory

	return app
}
