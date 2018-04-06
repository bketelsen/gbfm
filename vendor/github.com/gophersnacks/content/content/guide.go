package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Guide struct {
	item.Item

	GuideSlug   string   `json:"guide_slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	EmbedCode   string   `json:"embed_code"`
	Body        string   `json:"body"`
	Pro         bool     `json:"pro"`
	Keywords    []string `json:"keywords"`
	Author      string   `json:"author"`
}

// MarshalEditor writes a buffer of html to edit a Guide within the CMS
// and implements editor.Editable
func (g *Guide) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(g,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Guide field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("GuideSlug", g, map[string]string{
				"label":       "GuideSlug",
				"type":        "text",
				"placeholder": "Enter the GuideSlug here",
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
			View: editor.Textarea("Description", g, map[string]string{
				"label":       "Description - Markdown",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.File("Thumbnail", g, map[string]string{
				"label":       "Thumbnail",
				"placeholder": "Upload the Thumbnail here",
			}),
		},
		editor.Field{
			View: editor.Input("EmbedCode", g, map[string]string{
				"label":       "EmbedCode",
				"type":        "text",
				"placeholder": "Enter the EmbedCode here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Body", g, map[string]string{
				"label":       "Body - Markdown",
				"placeholder": "Enter the Body here",
			}),
		},
		editor.Field{
			View: editor.Input("Pro", g, map[string]string{
				"label":       "Pro",
				"type":        "text",
				"placeholder": "Enter the Pro here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Keywords", g, map[string]string{
				"label":       "Keywords",
				"type":        "text",
				"placeholder": "Enter the Keywords here",
			}),
		},
		editor.Field{
			View: reference.Select("Author", g, map[string]string{
				"label": "Author",
			},
				"Author",
				`{{ .name }} `,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Guide editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Guide"] = func() interface{} { return new(Guide) }
}

// String defines how a Guide is printed. Update it using more descriptive
// fields from the Guide struct type
func (g *Guide) String() string {
	return fmt.Sprintf("Guide: %s", g.UUID)
}
