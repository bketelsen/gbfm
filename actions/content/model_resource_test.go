package content

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/uuid"
	"github.com/gophersnacks/gbfm/models"
	"github.com/jinzhu/inflection"
)

func (as ActionSuite) TestModelList() {
	r, db := as.Require(), as.DB
	for modelName := range templateRegistry {
		r.NoError(as.login())
		as.T().Logf("model %s", modelName)
		singleModel, err := models.SampleFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))
		res := as.HTML("/admin/%s", inflection.Plural(modelName)).Get()
		r.Equal(http.StatusOK, res.Code)
	}
}

func (as ActionSuite) TestModelNew() {
	r := as.Require()
	for modelName := range templateRegistry {
		r.NoError(as.login())
		as.T().Logf("model %s", modelName)
		res := as.HTML("/admin/%s/new", modelName).Get()
		r.Equal(http.StatusOK, res.Code)
		res = as.HTML("/admin/%s/new)", inflection.Plural(modelName)).Get()
		r.Equal(http.StatusOK, res.Code)
	}
}

func (as ActionSuite) TestModelShow() {
	r, db := as.Require(), as.DB
	r.NoError(as.login())
	for modelName := range templateRegistry {
		as.T().Logf("model %s", modelName)
		singleModel, err := models.SampleFromRegistry(modelName)
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
		as.T().Logf("model %s", modelName)
		// create a new model
		singleModel, err := models.SampleFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))

		// send the DELETE to the destroy endpoint
		res := as.HTML("/admin/%s/%s", modelName, singleModel.GetID()).Delete()
		r.Equal(http.StatusFound, res.Code)
		r.Equal("/admin", res.Header().Get("location"))

		// make sure the model is gone from the DB
		r.NoError(as.modelIsGone(modelName, singleModel.GetID()))

		// now try making sure that the redir path param works
		singleModel, err = models.SampleFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.Create(singleModel))
		res = as.HTML(
			"/admin/%s/%s?redir_path=%s",
			modelName,
			singleModel.GetID(),
			"/otherPath",
		).Delete()
		r.Equal(http.StatusFound, res.Code)
		r.Equal("/otherPath", res.Header().Get("location"))
	}
}

func (as ActionSuite) TestModelCreate() {
	r, db := as.Require(), as.DB
	for modelName := range templateRegistry {
		as.T().Logf("model %s", modelName)
		singleModel, err := models.SampleFromRegistry(modelName)
		r.NoError(err)

		// make sure the endpoint returned the redirect
		res := as.HTML("/admin/%s", modelName).Post(singleModel)
		r.Equal(http.StatusFound, res.Code)

		// look for the new model in the DB
		list, err := models.EmptyListFromRegistry(modelName)
		r.NoError(err)
		r.NoError(db.All(list))
		r.Equal(1, list.Len())

		// make sure the response redirected to the right place
		r.Equal(
			fmt.Sprintf("/admin/%s/%s", modelName, list.EltAt(0).GetID()),
			res.Header().Get("Location"),
		)
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
