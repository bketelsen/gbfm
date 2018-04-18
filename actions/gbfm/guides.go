package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

func GuideList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	gList := new([]models.Guide)
	if err := tx.Eager().All(gList); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("guides", gList)
	return c.Render(200, r.HTML("guides/album.html"))
}

// SeriesShow gets the data for one Series.
func GuideShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	slug := c.Param("name")
	if slug == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	g := new(models.Guide)
	if err := tx.Eager().Where("slug = ?", slug).First(g); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("guides", g)
	return c.Render(200, r.HTML("guides/show.html"))
}
