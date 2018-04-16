package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// AuthorList gets all Authors. This function is mapped to the path
// GET /authors
func AuthorList(c buffalo.Context) error {
	list := new([]models.Author)
	if err := models.DB.Eager().All(list); err != nil {
		return c.Error(500, err)
	}
	c.Set("authors", list)
	return c.Render(200, r.HTML("authors/index.html"))
}

// AuthorShow gets the data for one Author. This function is mapped to
// the path GET /authors/{name} where name is the slug
func AuthorShow(c buffalo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	author := new(models.Author)
	if err := models.DB.Where("name = ?", name).First(author); err != nil {
		return c.Error(500, err)
	}
	c.Set("author", author)
	return c.Render(200, r.HTML("authors/show.html"))
}
