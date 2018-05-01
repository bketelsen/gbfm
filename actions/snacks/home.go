package snacks

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/ulule/paging"
)

// homeHandler is a default handler to serve up a home page
func homeHandler(c buffalo.Context) error {
	var snacks []models.Snack

	store, err := paging.NewGORMStore(models.DB, &snacks)
	if err != nil {
		log.Fatal(err)
	}

	options := paging.NewOptions()

	paginator, err := paging.NewOffsetPaginator(store, c.Request(), options)

	if err != nil {
		log.Fatal(err)
	}
	err = paginator.Page()
	if err != nil {
		log.Fatal(err)
	}
	//err := models.DB.Preload("Authors").Preload("Topics").Find(&snacks).Error
	//if err != nil {
	//	return c.Error(http.StatusInternalServerError, err)

	//}
	c.Set("snacks", snacks)
	c.Set("page", paginator)
	return c.Render(200, r.HTML("snacks/index.html"))
}
