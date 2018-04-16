package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/actions/content"
)

func init() {
	a, _ := content.App()
	buffalo.Grifts(a)
}
