package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/actions/gbfm"
)

func init() {
	buffalo.Grifts(gbfm.App())
}
