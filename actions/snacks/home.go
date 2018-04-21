package snacks

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
)

// homeHandler is a default handler to serve up a home page
func homeHandler(c buffalo.Context) error {
	var snacks []models.Snack
	err := models.DB.Preload("Authors").Preload("Topics").Find(&snacks).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)

	}
	c.Set("snacks", snacks)
	return c.Render(200, r.HTML("snacks/index.html"))
}
