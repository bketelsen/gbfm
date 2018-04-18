package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Episode represents an episode
type Episode struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Slug      string    `json:"slug" db:"slug"`

	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description" db:"description"`
	Markdown     string   `json:"markdown" db:"markdown"`
	ThumbnailURL string   `json:"thumbnail_url" db:"thumbnail_url"`
	EmbedCode    string   `json:"embed_code" db:"embed_code"`
	Body         string   `json:"body" db:"body"`
	Pro          bool     `json:"pro" db:"pro"`
	Repo         string   `json:"repo" db:"repo"`
	Topics       []Topic  `json:"topics" many_to_many:"episodes_topics"`
	Authors      []Author `json:"authors" many_to_many:"episodes_authors"`
	Series       []Series `json:"series_ids" many_to_many:"episodes_series"`
}

func init() {
	registry["episode"] = &registryFuncs{
		empty: func() IDer { return new(Episode) },
		list:  func() interface{} { return new([]Episode) },
		sample: func() IDer {
			return &Episode{
				Slug:         namer.NameSep("-"),
				Title:        namer.Name(),
				Description:  namer.Name(),
				ThumbnailURL: namer.NameSep("-"),
				EmbedCode:    namer.NameSep("-"),
				Body:         namer.Name(),
				Pro:          true,
				Repo:         namer.NameSep("-"),
			}
		},
	}
}

// GetID implements IDer
func (a Episode) GetID() uuid.UUID {
	return a.ID
}

// GetCreatedAt implements Core
func (a Episode) GetCreatedAt() time.Time {
	return a.CreatedAt
}

// GetUpdatedAt implements Core
func (a Episode) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

// GetSlug implements Slugger
func (a Episode) GetSlug() string {
	return a.Slug
}
