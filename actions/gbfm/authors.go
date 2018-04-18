package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// AuthorList gets all Episodes.
// GET /author
func AuthorList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	aList := new([]models.Author)
	if err := tx.Eager().All(aList); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("authors", aList)
	return c.Render(200, r.HTML("authors/album.html"))
}

// AuthorShow gets the data for one Author.
func AuthorShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	slug := c.Param("name")
	if slug == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	a := new(models.Author)
	if err := tx.Eager().Where("slug = ?", slug).First(a); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("author", a)
	return c.Render(200, r.HTML("authors/show.html"))
}
