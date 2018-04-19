package snacks

import (
	"net/http"

	"github.com/gophersnacks/gbfm/models"
)

func (as ActionSuite) TestSnackHandler() {
	const slug = "test-snack"
	r, db := as.Require(), as.DB

	// add a snack to the DB
	snack := &models.Snack{
		Slug:      slug,
		Title:     namer.Name(),
		Sponsored: true,
		URL:       namer.NameSep("-"),
		Summary:   namer.Name(),
		Comment:   namer.Name(),
		EmbedCode: namer.NameSep("-"),
	}
	r.NoError(db.Create(snack))
	res := as.HTML("/snack/%s", slug).Get()
	r.Equal(http.StatusOK, res.Code)
}
