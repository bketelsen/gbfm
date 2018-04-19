package render

import (
	"github.com/gobuffalo/buffalo/render"
)

// New returns a new render engine
func New(layout string, helpers render.Helpers) *render.Engine {
	return render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: layout,

		// Box containing all of the templates:
		TemplatesBox: TemplatesBox,
		AssetsBox:    AssetsBox,

		// Add template helpers here:
		Helpers: helpers,
	})
}
