package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// AuthorList gets all Authors. This function is mapped to the path
// GET /authors
func AuthorList(c buffalo.Context) error {
	aa, err := models.GetAuthorList()
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("authors", aa)
	return c.Render(200, r.HTML("authors/index.html"))
}

// AuthorShow gets the data for one Author. This function is mapped to
// the path GET /authors/{name} where name is the slug
func AuthorShow(c buffalo.Context) error {
	slug := c.Param("name")
	if slug == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	author, err := models.GetAuthorBySlug(slug)
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("author", author)
	return c.Render(200, r.HTML("authors/show.html"))
}
