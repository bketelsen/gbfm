package models

import (
	"fmt"
	"time"

	"github.com/bosssauce/reference"

	"github.com/gobuffalo/uuid"
	"github.com/gophersnacks/gbfm/pkg/management/editor"
	"github.com/gophersnacks/gbfm/pkg/system/item"
	suuid "github.com/satori/go.uuid"
)

func init() {
	item.Types["Episode"] = func() interface{} { return NewEpisode() }
}

// Episode represents an episode
type Episode struct {
	*ponzuSluggable
	*ponzuPartialIdentifiable
	*ponzuSearchable

	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Thumbnail   string    `json:"thumbnail" db:"thumbnail"`
	EmbedCode   string    `json:"embed_code" db:"embed_code"`
	Body        string    `json:"body" db:"body"`
	Pro         bool      `json:"pro" db:"pro"`
	Repo        string    `json:"repo" db:"repo"`
	Keywords    []string  `json:"keywords" db:"keywords"`
	Authors     []string  `json:"authors" db:"authors"`
}

// NewEpisode creates a new Episode. Use this instead of calling new(Episode)
func NewEpisode() *Episode {
	return &Episode{
		ponzuSluggable:           new(ponzuSluggable),
		ponzuPartialIdentifiable: new(ponzuPartialIdentifiable),
		ponzuSearchable:          new(ponzuSearchable),
	}
}

// UniqueID implements part of
// https://godoc.org/github.com/ponzu-cms/ponzu/system/item#Identifiable
func (e *Episode) UniqueID() suuid.UUID {
	return buffaloUUIDToSatori(e.ID)
}

// MarshalEditor writes a buffer of html to edit a Episode within the CMS
// and implements editor.Editable
func (e *Episode) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(e,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Episode field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Slug", e, map[string]string{
				"label":       "Slug",
				"type":        "text",
				"placeholder": "Enter the Slug here",
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

// String defines how a Episode is printed. Update it using more descriptive
// fields from the Episode struct type
func (e *Episode) String() string {
	return e.Slug
}
