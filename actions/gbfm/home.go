package gbfm

import "github.com/gobuffalo/buffalo"

// homeHandler is a default handler to serve up a home page
func homeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
