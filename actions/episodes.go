package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// EpisodeList gets all Episodes. This function is mapped to the path
// GET /episodes
func EpisodeList(c buffalo.Context) error {
	ee, err := models.GetEpisodeList()
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("episodes", ee)
	return c.Render(200, r.HTML("episodes/album.html"))
}

// EpisodeShow gets the data for one Episode. This function is mapped to
// the path GET /authors/{name} where name is the slug
func EpisodeShow(c buffalo.Context) error {
	slug := c.Param("name")
	if slug == "" {
		return c.Error(404, errors.New("Not Found"))
	}
	episode, err := models.GetFullEpisodeBySlug(slug)
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("episode", episode)
	return c.Render(200, r.HTML("episodes/show.html"))
}
