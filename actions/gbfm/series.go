package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Series)
// DB Table: Plural (series)
// Resource: Plural (Series)
// Path: Plural (/series)
// View Template Folder: Plural (/templates/series/)

// SeriesResource is the resource for the Series model
type SeriesResource struct {
	buffalo.BaseResource
}

// List gets all Series. This function is mapped to the path
// GET /series
func (v SeriesResource) List(c buffalo.Context) error {
	seriesList := []models.Series{}
	q := models.DB.Eager().PaginateFromParams(c.Request().URL.Query())
	if err := q.All(&seriesList); err != nil {
		c.Error(http.StatusInternalServerError, err)
	}
	c.Set("series", seriesList)
	c.Set("pagination", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("series/index.html"))
}

// Show gets the data for one Series. This function is mapped to
// the path GET /series/{series_id}
func (v SeriesResource) Show(c buffalo.Context) error {
	seriesID := c.Param("series_id")
	series := &models.Series{}
	if err := models.DB.Eager().Find(series, seriesID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("series", series)
	return c.Render(http.StatusOK, r.HTML("series/show.html"))
}
