package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
