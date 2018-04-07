package renderengine

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gophersnacks/gbfm/models"
)

// New returns a new render engine
func New(layout string) *render.Engine {
	return render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: layout,

		// Box containing all of the templates:
		TemplatesBox: TemplatesBox,
		AssetsBox:    AssetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"imageTag": func(source string) string {
				return models.BaseURL + source
			},
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
		},
	})
}
