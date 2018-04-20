package models

import (
	"fmt"
	"time"

	"github.com/gobuffalo/uuid"
)

type Series struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Slug         string    `json:"slug" db:"slug"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	ThumbnailURL string    `json:"thumbnail_url" db:"thumbnail_url"`
	Body         string    `json:"body" db:"body"`
	Pro          bool      `json:"pro" db:"pro"`
	Topics       []Topic   `json:"topics" many_to_many:"series_topics"`
	Authors      []Author  `json:"authors" many_to_many:"series_authors"`
	Episodes     []Episode `json:"episodes" many_to_many:"series_episodes" order_by:"created_at desc"`
}

func init() {
	registry["series"] = &registryFuncs{
		empty: func() IDer { return new(Series) },
		list:  func() Lister { return new(SeriesList) },
		sample: func() IDer {
			return &Series{
				Slug:         namer.NameSep("-"),
				Title:        namer.Name(),
				Description:  namer.Name(),
				ThumbnailURL: namer.NameSep("-"),
				Body:         namer.Name(),
				Pro:          true,
			}
		},
	}
}

func (a Series) String() string {
	return fmt.Sprintf("[%s] %s (%s)", a.ID, a.Title, a.Slug)
}

// GetID implements IDer
func (a Series) GetID() uuid.UUID {
	return a.ID
}

// GetCreatedAt implements Core
func (a Series) GetCreatedAt() time.Time {
	return a.CreatedAt
}

// GetUpdatedAt implements Core
func (a Series) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

// GetSlug implements Slugger
func (a Series) GetSlug() string {
	return a.Slug
}

// ModelName implements ModelNamer
func (Series) ModelName() string {
	return "Series"
}

// SeriesList is a list of Series models. It implements Lister
type SeriesList []*Series

// Len implements Lister
func (s SeriesList) Len() int {
	return len(s)
}

// EltAt implements Lister
func (s SeriesList) EltAt(i int) IDer {
	if i < len(s) {
		return s[i]
	}
	return nil
}
