package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// AuthorList gets all Episodes.
// GET /author
func AuthorList(c buffalo.Context) error {
	var authors []models.Author
	if err := models.GORM.Find(&authors).Error; err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("authors", authors)
	return c.Render(200, r.HTML("authors/album.html"))
}

// AuthorShow gets the data for one Author.
func AuthorShow(c buffalo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.Error(http.StatusBadRequest, errors.New("no epside slug found"))
	}
	id, err := idFromSlug(name)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	var author models.Author
	err = models.GORM.Where(string(id)).First(&author).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Set("author", author)
	return c.Render(200, r.HTML("authors/show.html"))
}
