package gbfm

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

func GuideList(c buffalo.Context) error {
	var guides []models.Guide
	err := models.GORM.Preload("Authors").Preload("Topics").Find(&guides).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)

	}
	c.Set("guides", guides)
	return c.Render(200, r.HTML("guides/album.html"))
}

func GuideShow(c buffalo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.Error(http.StatusBadRequest, errors.New("no epside slug found"))
	}
	id, err := idFromSlug(name)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	var guide models.Guide
	err = models.GORM.Preload("Topics").Preload("Authors").Where(id).First(&guide).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Set("guide", guide)
	return c.Render(200, r.HTML("guides/show.html"))
}
