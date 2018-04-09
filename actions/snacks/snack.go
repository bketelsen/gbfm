package snacks

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func snackHandler(c buffalo.Context) error {
	slug := c.Param("slug")
	c.Set("slug", slug)
	return c.Render(http.StatusOK, r.HTML("snack.html"))
}
