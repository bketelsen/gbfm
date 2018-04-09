package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Snack is an individual news snippet
//
// TODO: field(s) to indicate an embedded GIFM/GBFM episode
type Snack struct {
	item.Item

	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Sponsored bool     `json:"sponsored"`
	URL       string   `json:"url"`
	Summary   string   `json:"summary"`
	Comment   string   `json:"comment"`
	SnackSlug string   `json:"snack_slug"`
	Topics    []string `json:"topics"`
	Authors   []string `json:"authors"`
}

// MarshalEditor writes a buffer of html to edit a Snack within the CMS
// and implements editor.Editable
func (s *Snack) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(s,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Snack field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", s, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Author", s, map[string]string{
				"label":       "Author",
				"type":        "text",
				"placeholder": "Enter the Author here",
			}),
		},
		editor.Field{
			View: editor.Input("Sponsored", s, map[string]string{
				"label":       "Sponsored",
				"type":        "text",
				"placeholder": "Enter the Sponsored here",
			}),
		},
		editor.Field{
			View: editor.Input("URL", s, map[string]string{
				"label":       "URL",
				"type":        "text",
				"placeholder": "Enter the URL here",
			}),
		},
		editor.Field{
			View: editor.Input("Summary", s, map[string]string{
				"label":       "Summary",
				"type":        "text",
				"placeholder": "Enter the Summary here",
			}),
		},
		editor.Field{
			View: editor.Input("Comment", s, map[string]string{
				"label":       "Comment",
				"type":        "text",
				"placeholder": "Enter the Comment here",
			}),
		},
		editor.Field{
			View: editor.Input("SnackSlug", s, map[string]string{
				"label":       "SnackSlug",
				"type":        "text",
				"placeholder": "Enter the SnackSlug here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Topics", s, map[string]string{
				"label":       "Topics",
				"type":        "text",
				"placeholder": "Enter a topic here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Authors", s, map[string]string{
				"label": "Author",
			},
				"Author",
				"{{.name}}",
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Snack editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Snack"] = func() interface{} { return new(Snack) }
}

// String defines how a Snack is printed. Update it using more descriptive
// fields from the Snack struct type
func (s *Snack) String() string {
	return s.SnackSlug
}
