package content

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func adaptHandler(hdl http.HandlerFunc) buffalo.Handler {
	return func(c buffalo.Context) error {
		hdl(c.Response(), c.Request())
		return nil
	}
}
