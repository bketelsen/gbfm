package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/actions"
)

func init() {
	a, _ := actions.ContentApp()
	buffalo.Grifts(a)
}
