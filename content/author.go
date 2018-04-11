package content

import (
	"fmt"

	"github.com/gophersnacks/gbfm/pkg/management/editor"
	"github.com/gophersnacks/gbfm/pkg/system/item"
)

type Author struct {
	item.Item

	Name        string `json:"name"`
	Description string `json:"description"`
	AuthorSlug  string `json:"author_slug"`
	Photo       string `json:"photo"`
}

func (a *Author) String() string {
	return a.AuthorSlug
}

// MarshalEditor writes a buffer of html to edit a Author within the CMS
// and implements editor.Editable
func (a *Author) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Author field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", a, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", a, map[string]string{
				"label":       "Description - Markdown",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("AuthorSlug", a, map[string]string{
				"label":       "AuthorSlug",
				"type":        "text",
				"placeholder": "Enter the AuthorSlug here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", a, map[string]string{
				"label":       "Photo",
				"placeholder": "Upload the Photo here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Author editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Author"] = func() interface{} { return new(Author) }
}
