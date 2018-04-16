package gbfm

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/pkg/render"
)

var r = render.New("gbfm/application.html")

// AddRoutes adds routes for the Go Beyond Five Minutes site
func AddRoutes(app *buffalo.App) {
	app.GET("/", homeHandler)
	app.Resource("/series", SeriesResource{})
	app.Resource("/guides", GuidesResource{})
	app.GET("/episodes", EpisodeList)
	app.GET("/episodes/{name}", EpisodeShow)
}
