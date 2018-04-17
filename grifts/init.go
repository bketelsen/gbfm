package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/actions/content"
)

func init() {
	buffalo.Grifts(content.App())
}
