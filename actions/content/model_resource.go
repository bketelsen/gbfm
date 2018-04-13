package content

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
)

type modelResource struct{}

func (m *modelResource) List(c buffalo.Context) error {
	modelName := c.Param("model_name")
	tpls, err := getTemplateNames(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	list, err := models.EmptyListFromRegistry(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	if err := models.DB.All(&list); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("models", list)
	return c.Render(http.StatusOK, r.HTML(tpls.Index))
}
func (m *modelResource) Show(c buffalo.Context) error {
	modelName := c.Param("model_name")
	tpls, err := getTemplateNames(modelName)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	id := c.Param("admin_model_id")
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

func (m *modelResource) New(buffalo.Context) error     { return nil }
func (m *modelResource) Create(buffalo.Context) error  { return nil }
func (m *modelResource) Edit(buffalo.Context) error    { return nil }
func (m *modelResource) Update(buffalo.Context) error  { return nil }
func (m *modelResource) Destroy(buffalo.Context) error { return nil }
