package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// SeriesList gets all Episodes.
// GET /series
func SeriesList(c buffalo.Context) error {
	var ss []models.Series
	err := models.DB.Preload("Authors").Preload("Topics").Find(&ss).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)

	}
	c.Set("series", ss)
	return c.Render(200, r.HTML("series/album.html"))

}

// SeriesShow gets the data for one Series.
func SeriesShow(c buffalo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.Error(http.StatusBadRequest, errors.New("no epside slug found"))
	}
	id, err := idFromSlug(name)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	var s models.Series
	err = models.DB.Preload("Topics").Preload("Episodes").Preload("Authors").Where(id).First(&s).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Set("series", s)
	return c.Render(200, r.HTML("series/show.html"))
}
