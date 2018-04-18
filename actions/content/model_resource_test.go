package content

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/uuid"
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

func (as ActionSuite) TestModelShow() {
	r, db := as.Require(), as.DB
	for modelName := range templateRegistry {
		singleModel, err := models.EmptyFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))
		res := as.HTML("/admin/%s/%s", modelName, singleModel.GetID()).Get()
		r.Equal(http.StatusOK, res.Code)
	}
}

func (as ActionSuite) TestModelDestroy() {
	r, db := as.Require(), as.DB
	r.NoError(as.login())
	for modelName := range templateRegistry {
		// create a new model
		singleModel, err := models.EmptyFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))

		// send the DELETE to the destroy endpoint
		res := as.HTML("/admin/%s/%s", modelName, singleModel.GetID()).Delete()
		r.Equal(http.StatusFound, res.Code)
		r.Equal("/admin", res.Header().Get("location"))

		// make sure the model is gone from the DB
		r.NoError(as.modelIsGone(modelName, singleModel.GetID()))

		// now try making sure that the redir path param works
		singleModel, err = models.EmptyFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))
		res = as.HTML(
			"/admin/%s/%s?redir_path=%s",
			modelName,
			singleModel.GetID(),
			"/otherPath",
		).Delete()
		r.Equal(http.StatusFound, res.Code)
		r.Equal("/other_path", res.Header().Get("location"))
	}
}

func (as ActionSuite) modelIsGone(modelName string, id uuid.UUID) error {
	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return err
	}
	if err := as.DB.Find(empty, id); err == nil {
		return fmt.Errorf("expected model %s/%s to be missing", modelName, id)
	}
	return nil
}
