package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func homeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("root.html"))
}
