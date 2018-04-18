package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Guide is a guide
type Guide struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Slug         string    `json:"slug" db:"slug"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Markdown     string    `json:"markdown" db:"markdown"`
	ThumbnailURL string    `json:"thumbnail_url" db:"thumbnail_url"`
	EmbedCode    string    `json:"embed_code" db:"embed_code"`
	Body         string    `json:"body" db:"body"`
	Pro          bool      `json:"pro" db:"pro"`
	Topics       []Topic   `json:"topics" many_to_many:"guides_topics"`
	Authors      []Author  `json:"authors" many_to_many:"guides_authors"`
}

func init() {
	registry["guide"] = &registryFuncs{
		empty: func() IDer { return new(Guide) },
		sample: func() IDer {
			return &Guide{
				Slug:         namer.NameSep("-"),
				Title:        namer.Name(),
				Description:  namer.Name(),
				Markdown:     namer.Name(),
				ThumbnailURL: namer.Name(),
				EmbedCode:    namer.Name(),
				Body:         namer.Name(),
				Pro:          true,
			}
		},
		list: func() interface{} { return new([]Guide) },
	}
}

func (a Guide) GetID() uuid.UUID {
	return a.ID
}

func (a Guide) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Guide) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Guide) GetSlug() string {
	return a.Slug
}
