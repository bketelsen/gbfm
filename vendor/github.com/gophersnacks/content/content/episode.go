package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Episode struct {
	item.Item

	EpisodeSlug string   `json:"episode_slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	EmbedCode   string   `json:"embed_code"`
	Body        string   `json:"body"`
	Pro         bool     `json:"pro"`
	Repo        string   `json:"repo"`
	Keywords    []string `json:"keywords"`
	Authors     []string `json:"authors"`
}

// MarshalEditor writes a buffer of html to edit a Episode within the CMS
// and implements editor.Editable
func (e *Episode) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(e,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Episode field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("EpisodeSlug", e, map[string]string{
				"label":       "EpisodeSlug",
				"type":        "text",
				"placeholder": "Enter the EpisodeSlug here",
			}),
		},
		editor.Field{
			View: editor.Input("Title", e, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Description", e, map[string]string{
				"label":       "Description - Markdown",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.File("Thumbnail", e, map[string]string{
				"label":       "Thumbnail",
				"placeholder": "Upload the Thumbnail here",
			}),
		},
		editor.Field{
			View: editor.Input("EmbedCode", e, map[string]string{
				"label":       "EmbedCode",
				"type":        "text",
				"placeholder": "Enter the EmbedCode here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Body", e, map[string]string{
				"label":       "Body - Markdown",
				"placeholder": "Enter the Body here",
			}),
		},
		editor.Field{
			View: editor.Input("Pro", e, map[string]string{
				"label":       "Pro",
				"type":        "text",
				"placeholder": "Enter the Pro here",
			}),
		},
		editor.Field{
			View: editor.Input("Repo", e, map[string]string{
				"label":       "Repo",
				"type":        "text",
				"placeholder": "Enter the Source Code Repository here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Keywords", e, map[string]string{
				"label":       "Keywords",
				"type":        "text",
				"placeholder": "Enter the Keywords here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Authors", e, map[string]string{
				"label": "Author",
			},
				"Author",
				"{{.name}}",
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Episode editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Episode"] = func() interface{} { return new(Episode) }
}

// String defines how a Episode is printed. Update it using more descriptive
// fields from the Episode struct type
func (e *Episode) String() string {
	return e.EpisodeSlug
}
