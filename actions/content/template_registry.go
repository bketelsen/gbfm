package content

import (
	"fmt"

	"github.com/jinzhu/inflection"
)

var templateRegistry = map[string]*templateNames{
	"author": {
		Show:  "authors/show.html",
		New:   "authors/new.html",
		Edit:  "authors/edit.html",
		Index: "authors/index.html",
	},
	"episode": {
		Show:  "episodes/show.html",
		New:   "episodes/new.html",
		Edit:  "episodes/edit.html",
		Index: "episodes/index.html",
	},
	"gifm": {
		Show:  "gifm/show.html",
		New:   "gifm/new.html",
		Edit:  "gifm/edit.html",
		Index: "gifm/index.html",
	},
	"guide": {
		Show:  "guides/show.html",
		New:   "guides/new.html",
		Edit:  "guides/edit.html",
		Index: "guides/index.html",
	},
	"series": {
		Show:  "series/show.html",
		New:   "series/new.html",
		Edit:  "series/edit.html",
		Index: "series/index.html",
	},
	"snack": {
		Show:  "snacks/show.html",
		New:   "snacks/new.html",
		Edit:  "snacks/edit.html",
		Index: "snacks/index.html",
	},
	"topic": {
		Show:  "topics/show.html",
		New:   "topics/new.html",
		Edit:  "topics/edit.html",
		Index: "topics/index.html",
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
