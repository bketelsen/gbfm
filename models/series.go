package models

import (
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
	ThumbnailURL string    `json:"thumbnail_url" db:"thumbnail_url`
	Body         string    `json:"body" db:"body"`
	Pro          bool      `json:"pro" db:"pro"`
	Topics       []Topic   `json:"topics" many_to_many:"series_topics"`
	Authors      []Author  `json:"authors" many_to_many:"series_authors"`
}

func init() {
	registry["series"] = func() (IDer, interface{}) {
		return new(Series), new([]Series)
	}
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
