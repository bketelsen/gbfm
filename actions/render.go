package actions

import (
	rend "github.com/gobuffalo/buffalo/render"
	"github.com/gophersnacks/gbfm/pkg/render"
)

var r = render.New("application.html", rend.Helpers{})
