package render

import (
	"github.com/gobuffalo/packr"
)

// AssetsBox is the packr box for public assets
var AssetsBox = packr.NewBox("../../public")

// TemplatesBox is the packr box for the templates dir
var TemplatesBox = packr.NewBox("../../templates")
