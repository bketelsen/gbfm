package snacks

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
)

// homeHandler is a default handler to serve up a home page
func homeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	snacks := new([]models.Snack)
	if err := tx.Eager().All(snacks); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("snacks", snacks)
	return c.Render(200, r.HTML("snacks/index.html"))
}
