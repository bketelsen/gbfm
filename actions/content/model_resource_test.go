package content

import (
	"net/http"

	"github.com/gophersnacks/gbfm/models"
)

func (as ActionSuite) TestModelList() {
	r, db := as.Require(), as.DB
	for modelName := range templateRegistry {
		singleModel, err := models.EmptyFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))
		res := as.HTML("/admin/%s", modelName).Get()
		r.Equal(http.StatusOK, res.Code)
	}
}
