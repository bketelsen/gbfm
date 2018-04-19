package gbfm

import (
	"github.com/gophersnacks/gbfm/models"
)

func (as ActionSuite) TestAuthorList() {
	const slug = "test-author"
	r, db := as.Require(), as.DB
	author := &models.Author{
		Slug:        slug,
		Name:        namer.Name(),
		Description: namer.Name(),
		Photo:       namer.Name(),
	}
	r.NoError(db.Create(author))

}
