package snacks

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/actions/renderengine"
)

var r = renderengine.New("../..")

// AddRoutes adds routes for the GopherSnacks site
func AddRoutes(app *buffalo.App) {
	app.GET("/", homeHandler)
}
