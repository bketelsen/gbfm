package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// EpisodeList gets all Episodes. This function is mapped to the path
// GET /episodes
func EpisodeList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	epList := new([]models.Episode)
	if err := tx.Eager().All(epList); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("episodes", epList)
	return c.Render(200, r.HTML("episodes/album.html"))
}

// EpisodeShow gets the data for one Episode. This function is mapped to
// the path GET /authors/{name} where name is the slug
func EpisodeShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	slug := c.Param("name")
	if slug == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	episode := new(models.Episode)
	if err := tx.Eager().Where("slug = ?", slug).First(episode); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("episode", episode)
	return c.Render(200, r.HTML("episodes/show.html"))
}
