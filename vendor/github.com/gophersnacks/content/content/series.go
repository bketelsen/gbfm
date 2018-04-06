package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Series struct {
	item.Item

	SeriesSlug  string   `json:"series_slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Body        string   `json:"body"`
	Pro         bool     `json:"pro"`
	Keywords    []string `json:"keywords"`
	Episodes    []string `json:"episodes"`
}

// MarshalEditor writes a buffer of html to edit a Series within the CMS
// and implements editor.Editable
func (s *Series) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(s,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Series field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("SeriesSlug", s, map[string]string{
				"label":       "SeriesSlug",
				"type":        "text",
				"placeholder": "Enter the SeriesSlug here",
			}),
		},
		editor.Field{
			View: editor.Input("Title", s, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", s, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.File("Thumbnail", s, map[string]string{
				"label":       "Thumbnail",
				"placeholder": "Upload the Thumbnail here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Body", s, map[string]string{
				"label":       "Body",
				"placeholder": "Enter the Body here",
			}),
		},
		editor.Field{
			View: editor.Input("Pro", s, map[string]string{
				"label":       "Pro",
				"type":        "text",
				"placeholder": "Enter the Pro here",
			}),
		},

		editor.Field{
			View: reference.SelectRepeater("Episodes", s, map[string]string{
				"label": "Episode",
			},
				"Episode",
				"{{.title}}",
			),
		},
		editor.Field{
			View: editor.InputRepeater("Keywords", s, map[string]string{
				"label":       "Keywords",
				"type":        "text",
				"placeholder": "Enter the Keywords here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Series editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Series"] = func() interface{} { return new(Series) }
}

// String defines how a Series is printed. Update it using more descriptive
// fields from the Series struct type
func (s *Series) String() string {
	return s.SeriesSlug
}
