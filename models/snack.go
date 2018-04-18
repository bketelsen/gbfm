package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Snack is a snack - a small piece of news
type Snack struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Slug      string    `json:"slug" db:"slug"`

	Title     string   `json:"title" db:"title"`
	Sponsored bool     `json:"sponsored" db:"sponsored"`
	URL       string   `json:"url" db:"url"`
	Summary   string   `json:"summary" db:"summary"`
	Comment   string   `json:"comment" db:"summary"`
	Topics    []Topic  `json:"topics" many_to_many:"topics_snacks"`
	Authors   []Author `json:"authors" many_to_many:"authors_snacks"`
}

func init() {
	registry["snack"] = &registryFuncs{
		list:  func() interface{} { return new([]Snack) },
		empty: func() IDer { return new(Snack) },
		sample: func() IDer {
			return &Snack{
				Slug:      namer.NameSep("-"),
				Title:     namer.Name(),
				Sponsored: true,
				URL:       namer.NameSep("-"),
				Summary:   namer.Name(),
				Comment:   namer.Name(),
			}
		},
	}
}

func (a Snack) GetID() uuid.UUID {
	return a.ID
}

func (a Snack) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Snack) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Snack) GetSlug() string {
	return a.Slug
}
