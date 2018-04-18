package content

import (
	"errors"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
)

type modelResource struct{}

// /admin/{model_name}
func (m *modelResource) List(c buffalo.Context) error {
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	tpls, err := getTemplateNames(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	list, err := models.EmptyListFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := models.DB.All(list); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("models", list)
	return c.Render(http.StatusOK, r.HTML(tpls.Index))
}

// /admin/{model_name}/{admin_model_id}
func (m *modelResource) Show(c buffalo.Context) error {
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	id, err := getModelID(c)
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
	if err := models.DB.Where("id = ?", id).First(single); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("model", single)
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

	c.Set("model_name", modelName)
	return c.Render(http.StatusOK, r.HTML(templateNames.New))
}

// POST /admin/{model_name}
func (m *modelResource) Create(c buffalo.Context) error {
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

	return c.Redirect(http.StatusFound, "/admin/%s/%s", modelName, empty.GetID())
}

// GET /admin/{model_name}/{admin_model_id}/edit
func (m *modelResource) Edit(c buffalo.Context) error {
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	modelID, err := getModelID(c)
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
	if err := models.DB.Find(empty, modelID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("model", empty)
	return c.Render(http.StatusOK, r.HTML(templateNames.Edit))
}

func (m *modelResource) Update(buffalo.Context) error { return nil }

// DELETE /admin/{model_name}/{admin_model_id}
func (m *modelResource) Destroy(c buffalo.Context) error {
	modelName, err := getModelName(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	modelID, err := getModelID(c)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	empty, err := models.EmptyFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := models.DB.Find(empty, modelID); err != nil {
		return c.Error(http.StatusNotAcceptable, err)
	}
	if err := models.DB.Destroy(empty); err != nil {
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
		return "", errors.New("model name not found")
	}
	return modelName, nil
}

func getModelID(c buffalo.Context) (string, error) {
	modelID := c.Param("admin_model_id")
	if modelID == "" {
		return "", errors.New("model ID not found")
	}
	return modelID, nil
}
