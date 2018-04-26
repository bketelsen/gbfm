package snacks

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
)

func topicHandler(c buffalo.Context) error {
	var topics []models.Topic
	err := models.DB.Find(&topics).Error
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)

	}
	c.Set("topics", topics)
	return c.Render(http.StatusOK, r.HTML("topics/index.html"))
}
