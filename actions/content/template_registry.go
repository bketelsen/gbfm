package content

import (
	"fmt"
)

var templateRegistry = map[string]*templateNames{
	"author": {
		Show:  "admin/author/show.html",
		New:   "admin/author/new.html",
		Edit:  "admin/author/edit.html",
		Index: "admin/author/index.html",
	},
	"episode": {
		Show:  "admin/episode/show.html",
		New:   "admin/episode/new.html",
		Edit:  "admin/episode/edit.html",
		Index: "admin/episode/index.html",
	},
	"gifm": {
		Show:  "admin/gifm/show.html",
		New:   "admin/gifm/new.html",
		Edit:  "admin/gifm/edit.html",
		Index: "admin/gifm/index.html",
	},
	"guide": {
		Show:  "admin/guide/show.html",
		New:   "admin/guide/new.html",
		Edit:  "admin/guide/edit.html",
		Index: "admin/guide/index.html",
	},
	"series": {
		Show:  "admin/series/show.html",
		New:   "admin/series/new.html",
		Edit:  "admin/series/edit.html",
		Index: "admin/series/index.html",
	},
	"snack": {
		Show:  "admin/snack/show.html",
		New:   "admin/snack/new.html",
		Edit:  "admin/snack/edit.html",
		Index: "admin/snack/index.html",
	},
	"topic": {
		Show:  "admin/topic/show.html",
		New:   "admin/topic/new.html",
		Edit:  "admin/topic/edit.html",
		Index: "admin/topic/index.html",
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
		return nil, fmt.Errorf("unknown model %s", s)
	}
	return tn, nil
}
