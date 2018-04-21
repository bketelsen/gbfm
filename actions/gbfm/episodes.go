package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// EpisodeList gets all Episodes. This function is mapped to the path
// GET /episodes
func EpisodeList(c buffalo.Context) error {
	var episodes []models.Episode
	err := models.DB.Preload("Authors").Preload("Topics").Find(&episodes).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)

	}
	c.Set("episodes", episodes)
	return c.Render(200, r.HTML("episodes/album.html"))
}

// EpisodeShow gets the data for one Episode. This function is mapped to
// the path GET /authors/{name} where name is the slug
func EpisodeShow(c buffalo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.Error(http.StatusBadRequest, errors.New("no epside slug found"))
	}
	id, err := idFromSlug(name)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	var episode models.Episode
	err = models.DB.Preload("Topics").Preload("Authors").Where(id).First(&episode).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Set("episode", episode)
	return c.Render(200, r.HTML("episodes/show.html"))
}
