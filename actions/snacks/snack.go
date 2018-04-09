package snacks

import (
	"errors"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func snackHandler(c buffalo.Context) error {
	slug := c.Param("snack_slug")
	if slug == "" {
		return c.Error(http.StatusBadRequest, errors.New("no snack slug found"))
	}
	c.Set("slug", slug)
	return c.Render(http.StatusOK, r.HTML("snacks/snack.html"))
}
