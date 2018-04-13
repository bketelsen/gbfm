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
}

// TemplateNames holds the template names for creating, editing and showing a model
type templateNames struct {
	Show  string
	New   string
	Edit  string
	Index string
}

func getTemplateNames(s string) (*templateNames, error) {
	tn, ok := templateRegistry[s]
	if !ok {
		return nil, fmt.Errorf("unknown model %s", s)
	}
	return tn, nil
}
