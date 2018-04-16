package snacks

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/pkg/render"
)

var r = render.New("snacks/application.html")

// AddRoutes adds routes for the GopherSnacks site
func AddRoutes(app *buffalo.App) {
	app.GET("/", homeHandler)
	app.GET("/snack/{snack_slug}", snackHandler)
}
