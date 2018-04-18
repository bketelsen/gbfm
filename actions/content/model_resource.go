package content

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
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
	tpls, err := getTemplateNames(modelName)
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
	c.Set(modelName, list)
	return c.Render(http.StatusOK, r.HTML(tpls.Index))
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

	tpls, err := getTemplateNames(modelName)
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
	c.Set(modelName, single)
	return c.Render(http.StatusOK, r.HTML(tpls.Show))
}

// GET /admin/{model_name}/new
func (m *modelResource) New(c buffalo.Context) error {
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	templateNames, err := getTemplateNames(modelName)
	if err != nil {
		// TODO: better error message
		return c.Error(http.StatusBadRequest, err)
	}

	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("model_name", modelName)
	c.Set(modelName, empty)
	return c.Render(http.StatusOK, r.HTML(templateNames.New))
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
		return c.Error(http.StatusBadRequest, err)
	}
	if err := c.Bind(empty); err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := tx.Create(empty); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusFound, "/admin/%s/%s", modelName, empty.GetID())
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
	templateNames, err := getTemplateNames(modelName)
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
	c.Set("model", empty)
	return c.Render(http.StatusOK, r.HTML(templateNames.Edit))
}

func (m *modelResource) Update(buffalo.Context) error { return nil }

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
	return c.Redirect(http.StatusFound, redirPath)
}

func getModelName(c buffalo.Context) (string, error) {
	modelName := c.Param("model_name")
	if modelName == "" {
		return "", fmt.Errorf("model name %s not found", modelName)
	}
	return modelName, nil
}

func getModelID(c buffalo.Context, name string) (string, error) {
	modelID := c.Param("model_id")
	if modelID == "" {
		return "", fmt.Errorf("model ID not found for %s", name)
	}
	return modelID, nil
}
