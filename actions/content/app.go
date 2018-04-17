package content

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gophersnacks/gbfm/models"
	"github.com/gophersnacks/gbfm/pkg/render"
	"github.com/gophersnacks/gbfm/pkg/web"
	"github.com/unrolled/secure"
)

var r = render.New("common_layout.html")

// App is where all routes and middleware for the admin interface are defined.
//
// The second parameter returned should be called by the caller in a defer
func App() *buffalo.App {
	app := buffalo.New(buffalo.Options{
		Addr:        "0.0.0.0:8080",
		Env:         web.ENV,
		SessionName: "_admin_session",
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
		"Gopher Enterprises",
		"admin/nav.html",
		"admin/footer.html",
	))

	// Wraps each request in a transaction.
	//  c.Value("tx").(*pop.PopTransaction)
	// Remove to disable this.
	app.Use(middleware.PopTransaction(models.DB))

	// Setup and use translations:
	app.Use(web.Translator.Middleware())

	app.Use(Auth)

	app.GET("/", homeHandler)
	app.GET("/admin", homeHandler)
	mResource := &modelResource{}

	login := &loginHandlers{
		successRedir: "/admin",
		showFormPath: "/admin/login",
	}
	app.GET("/admin/login", login.showForm)
	app.POST("/admin/login", login.try)
	app.Middleware.Skip(Auth, login.showForm, login.try)

	app.Resource("/admin/{model_name}", mResource)

	app.ServeFiles("/", render.AssetsBox) // serve files from the public directory

	return app
}
