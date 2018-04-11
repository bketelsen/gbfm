package content

import (
	"fmt"

	"github.com/gophersnacks/gbfm/pkg/management/editor"
	"github.com/gophersnacks/gbfm/pkg/system/item"
)

type Gifm struct {
	item.Item

	Slug        string `json:"slug"`
	Title       string `json:"title"`
	EmdedCode   string `json:"emded_code"`
	GithubLink  string `json:"github_link"`
	Sponsor     string `json:"sponsor"`
	Description string `json:"description"`
}

func (g *Gifm) String() string {
	return g.Slug
}

// MarshalEditor writes a buffer of html to edit a Gifm within the CMS
// and implements editor.Editable
func (g *Gifm) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(g,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Gifm field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:

		editor.Field{
			View: editor.Input("Slug", g, map[string]string{
				"label":       "Slug",
				"type":        "text",
				"placeholder": "Enter the Slug here",
			}),
		},
		editor.Field{
			View: editor.Input("Title", g, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Textarea("EmdedCode", g, map[string]string{
				"label":       "EmdedCode",
				"placeholder": "Enter the EmdedCode here",
			}),
		},
		editor.Field{
			View: editor.Input("GithubLink", g, map[string]string{
				"label":       "GithubLink",
				"type":        "text",
				"placeholder": "Enter the GithubLink here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Sponsor", g, map[string]string{
				"label":       "Sponsor",
				"placeholder": "Enter the Sponsor here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", g, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Gifm editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Gifm"] = func() interface{} { return new(Gifm) }
}
