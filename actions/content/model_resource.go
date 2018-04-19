package content

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/jinzhu/inflection"
)

type modelResource struct {
}

// /admin/{model_name}
func (m *modelResource) List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	modelName, err := getModelName(c)
	if err != nil {
		c.Logger().Errorf("model name %s not found", modelName)
		return c.Error(http.StatusBadRequest, err)
	}
	templateInfo, err := getTemplateNames(modelName)
	if err != nil {
		c.Logger().Errorf("template for model %s not found", modelName)
		return c.Error(http.StatusBadRequest, err)
	}
	list, err := models.EmptyListFromRegistry(modelName)
	if err != nil {
		c.Logger().Errorf("model registry lookup for %s", modelName)
		return c.Error(http.StatusBadRequest, err)
	}
	q := tx.Q().PaginateFromParams(c.Request().URL.Query())
	if err := q.All(list); err != nil {
		c.Logger().Errorf("fetching model list for %s", modelName)
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("pagination", q.Paginator)
	plural := inflection.Plural(modelName)
	c.Set(plural, list)
	return c.Render(http.StatusOK, r.HTML(templateInfo.Index))
}

// /admin/{model_name}/{admin_model_id}
func (m *modelResource) Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	id, err := getModelID(c, modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	templateInfo, err := getTemplateNames(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	single, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := tx.Where("id = ?", id).First(single); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Set("model_name", modelName)
	c.Set(modelName, single)
	return c.Render(http.StatusOK, r.HTML(templateInfo.Show))
}

// GET /admin/{model_name}s/new
func (m *modelResource) New(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	templateInfo, err := getTemplateNames(modelName)
	if err != nil {
		c.Logger().Errorf("getting template names for model %s", modelName)
		return c.Error(http.StatusBadRequest, err)
	}

	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		c.Logger().Errorf("getting empty model %s", modelName)
		return c.Error(http.StatusNotFound, err)
	}

	if err := templateInfo.fetchAdditionalModels(tx); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	templateInfo.populateAdditionalModels(c)

	c.Set(modelName, empty)
	// this allows URLs to have plural names (i.e. /admin/episodes/new)
	// but the template can still use the singular name (i.e. episode)
	c.Set(inflection.Singular(modelName), empty)
	return c.Render(http.StatusOK, r.HTML(templateInfo.New))
}

// POST /admin/{model_name}
func (m *modelResource) Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		c.Logger().Errorf("getting empty model %s from registry", modelName)
		return c.Error(http.StatusBadRequest, err)
	}
	if err := c.Bind(empty); err != nil {
		c.Logger().Errorf("binding to model %s", modelName)
		return c.Error(http.StatusBadRequest, err)
	}
	verrs, err := tx.Eager().ValidateAndCreate(empty)
	if verrs.HasAny() {
		c.Logger().Errorf("ValidateAndCreate on a new %s (%s)", modelName, verrs)
		return c.Error(http.StatusBadRequest, verrs)
	} else if err != nil {
		c.Logger().Errorf("creating a new %s (%s)", modelName, err)
		return c.Error(http.StatusInternalServerError, err)
	}

	singular := inflection.Singular(modelName)
	c.Set("model_name", modelName)
	c.Set(modelName, singular)
	return c.Redirect(http.StatusFound, "/admin/%s/%s", singular, empty.GetID())
}

// GET /admin/{model_name}/{admin_model_id}/edit
func (m *modelResource) Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	modelID, err := getModelID(c, modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	templateInfo, err := getTemplateNames(modelName)
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := tx.Find(empty, modelID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := templateInfo.fetchAdditionalModels(tx); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	templateInfo.populateAdditionalModels(c)

	c.Set("model_name", modelName)
	c.Set("model_name_plural", inflection.Plural(modelName))
	c.Set("model_id", modelID)
	c.Set(modelName, empty)
	return c.Render(http.StatusOK, r.HTML(templateInfo.Edit))
}

// PUT /admin/{model_name}/{model_id}
func (m *modelResource) Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	modelID, err := getModelID(c, modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	if err := tx.Find(empty, modelID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	if err := c.Bind(empty); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	verrs, err := tx.ValidateAndUpdate(empty)
	if verrs.HasAny() {
		return c.Error(http.StatusInternalServerError, verrs)
	} else if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusFound, "/admin/%s/%s", inflection.Plural(modelName), empty.GetID())
}

// DELETE /admin/{model_name}/{admin_model_id}
func (m *modelResource) Destroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	modelID, err := getModelID(c, modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := tx.Find(empty, modelID); err != nil {
		return c.Error(http.StatusNotAcceptable, err)
	}
	if err := tx.Destroy(empty); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	redirPath := "/admin"
	if c.Param("redir_path") != "" {
		redirPath = c.Param("redir_path")
	}
	c.Set("model_name", modelName)
	return c.Redirect(http.StatusFound, redirPath)
}

func getModelName(c buffalo.Context) (string, error) {
	modelName := c.Param("model_name")
	if modelName == "" {
		return "", fmt.Errorf("model name %s not found", modelName)
	}

	modelName = inflection.Singular(modelName)
	return modelName, nil
}

func getModelID(c buffalo.Context, name string) (string, error) {
	modelID := c.Param("model_id")
	if modelID == "" {
		return "", fmt.Errorf("model ID not found for %s", name)
	}
	return modelID, nil
}
