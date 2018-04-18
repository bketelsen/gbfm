package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// SeriesList gets all Episodes.
// GET /series
func SeriesList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	sList := new([]models.Series)
	if err := tx.Eager().All(sList); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("series", sList)
	return c.Render(200, r.HTML("series/album.html"))
}

// SeriesShow gets the data for one Series.
func SeriesShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	slug := c.Param("name")
	if slug == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	s := new(models.Series)
	if err := tx.Eager().Where("slug = ?", slug).First(s); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("series", s)
	return c.Render(200, r.HTML("series/show.html"))
}
