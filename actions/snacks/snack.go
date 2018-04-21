package snacks

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
)

func snackHandler(c buffalo.Context) error {
	slug := c.Param("snack_slug")
	if slug == "" {
		return c.Error(http.StatusBadRequest, errors.New("no snack slug found"))
	}
	id, err := idFromSlug(slug)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	var snack models.Snack
	err = models.DB.Preload("Topics").Preload("Authors").Where(id).First(&snack).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Set("snack", snack)
	return c.Render(http.StatusOK, r.HTML("snacks/snack.html"))
}

func idFromSlug(s string) (uint, error) {
	pieces := strings.Split(s, "-")
	if len(pieces) < 1 {
		return 0, errors.New("bad article")
	}
	id := pieces[0]
	uid, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("bad article ID")
	}
	return uint(uid), nil
}
