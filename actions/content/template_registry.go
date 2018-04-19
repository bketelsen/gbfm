package content

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/jinzhu/inflection"
)

var templateRegistry = map[string]*templateInfo{
	"author": {
		Show:             "admin/authors/show.html",
		New:              "admin/authors/new.html",
		Edit:             "admin/authors/edit.html",
		Index:            "admin/authors/index.html",
		AdditionalModels: map[string]func() models.Lister{},
	},
	"episode": {
		Show:  "admin/episodes/show.html",
		New:   "admin/episodes/new.html",
		Edit:  "admin/episodes/edit.html",
		Index: "admin/episodes/index.html",
		AdditionalModels: map[string]func() models.Lister{
			"topics":     newTopicsList,
			"authors":    newAuthorsList,
			"seriesList": newSeriesList,
		},
	},
	"gbfm": {
		Show:  "admin/gbfm/show.html",
		New:   "admin/gbfm/new.html",
		Edit:  "admin/gbfm/edit.html",
		Index: "admin/gbfm/index.html",
		AdditionalModels: map[string]func() models.Lister{
			"topics":  newTopicsList,
			"authors": newAuthorsList,
		},
	},
	"guide": {
		Show:  "admin/guides/show.html",
		New:   "admin/guides/new.html",
		Edit:  "admin/guides/edit.html",
		Index: "admin/guides/index.html",
		AdditionalModels: map[string]func() models.Lister{
			"topics":  newTopicsList,
			"authors": newAuthorsList,
		},
	},
	"image": {
		Show:             "admin/images/show.html",
		New:              "admin/images/new.html",
		Edit:             "admin/images/edit.html",
		Index:            "admin/images/index.html",
		AdditionalModels: map[string]func() models.Lister{},
	},
	"series": {
		Show:  "admin/series/show.html",
		New:   "admin/series/new.html",
		Edit:  "admin/series/edit.html",
		Index: "admin/series/index.html",
		AdditionalModels: map[string]func() models.Lister{
			"topics":  newTopicsList,
			"authors": newAuthorsList,
		},
	},
	"snack": {
		Show:  "admin/snacks/show.html",
		New:   "admin/snacks/new.html",
		Edit:  "admin/snacks/edit.html",
		Index: "admin/snacks/index.html",
		AdditionalModels: map[string]func() models.Lister{
			"topics":  newTopicsList,
			"authors": newAuthorsList,
		},
	},
	"topic": {
		Show:             "admin/topics/show.html",
		New:              "admin/topics/new.html",
		Edit:             "admin/topics/edit.html",
		Index:            "admin/topics/index.html",
		AdditionalModels: map[string]func() models.Lister{},
	},
}

// templateInfo holds the template names for creating, editing and showing a
// model, and also additional information that needs to be passed to the template
// for it to render properly
type templateInfo struct {
	Show  string
	New   string
	Edit  string
	Index string
	// the additional models that the handler must fetch and pass into the
	// template. The keys of this map are the names that must be passed
	// to the template
	AdditionalModels map[string]func() models.Lister
}

// fetchAdditionalModels uses tx to fetch every model.Lister in
// t.AdditionalModels. returns nil and an error if there was a failure fetching
// any of them, the full map otherwise. Use the return value of this method to
// pass into populateAdditionalModels
func (t templateInfo) fetchAdditionalModels(tx *pop.Connection) (map[string]models.Lister, error) {
	ret := map[string]models.Lister{}
	for name, listerFn := range t.AdditionalModels {
		lister := listerFn()
		if err := tx.All(lister); err != nil {
			return nil, err
		}
		ret[name] = lister
	}
	return ret, nil
}

// populateAdditionalModels sets each model list in t.AdditionalModels
// onto c. You should call t.fetchAdditionalModels before calling this.
func (t templateInfo) populateAdditionalModels(c buffalo.Context, m map[string]models.Lister) {
	for name, lister := range m {
		c.Set(name, lister)
	}
}

// TODO: deal with plurals
func getTemplateNames(s string) (*templateInfo, error) {
	tn, ok := templateRegistry[s]
	if !ok {
		singular := inflection.Singular(s)
		depluralized, depluralizedFound := templateRegistry[singular]
		if !depluralizedFound {
			return nil, fmt.Errorf("unknown model %s (singular: %s)", s, singular)
		}
		tn = depluralized
	}
	return tn, nil
}

func newTopicsList() models.Lister  { return &models.Topics{} }
func newAuthorsList() models.Lister { return &models.Authors{} }
func newSeriesList() models.Lister  { return &models.SeriesList{} }
