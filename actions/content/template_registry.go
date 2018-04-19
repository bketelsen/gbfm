package content

import (
	"fmt"

	"github.com/jinzhu/inflection"
)

var templateRegistry = map[string]*templateNames{
	"author": {
		Show:  "admin/authors/show.html",
		New:   "admin/authors/new.html",
		Edit:  "admin/authors/edit.html",
		Index: "admin/authors/index.html",
	},
	"episode": {
		Show:  "admin/episodes/show.html",
		New:   "admin/episodes/new.html",
		Edit:  "admin/episodes/edit.html",
		Index: "admin/episodes/index.html",
	},
	"gbfm": {
		Show:  "admin/gbfm/show.html",
		New:   "admin/gbfm/new.html",
		Edit:  "admin/gbfm/edit.html",
		Index: "admin/gbfm/index.html",
	},
	"guide": {
		Show:  "admin/guides/show.html",
		New:   "admin/guides/new.html",
		Edit:  "admin/guides/edit.html",
		Index: "admin/guides/index.html",
	},
	"image": {
		Show:  "admin/images/show.html",
		New:   "admin/images/new.html",
		Edit:  "admin/images/edit.html",
		Index: "admin/images/index.html",
	},
	"series": {
		Show:  "admin/series/show.html",
		New:   "admin/series/new.html",
		Edit:  "admin/series/edit.html",
		Index: "admin/series/index.html",
	},
	"snack": {
		Show:  "admin/snacks/show.html",
		New:   "admin/snacks/new.html",
		Edit:  "admin/snacks/edit.html",
		Index: "admin/snacks/index.html",
	},
	"topic": {
		Show:  "admin/topics/show.html",
		New:   "admin/topics/new.html",
		Edit:  "admin/topics/edit.html",
		Index: "admin/topics/index.html",
	},
}

// TemplateNames holds the template names for creating, editing and showing a model
type templateNames struct {
	Show  string
	New   string
	Edit  string
	Index string
}

// TODO: deal with plurals
func getTemplateNames(s string) (*templateNames, error) {
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
