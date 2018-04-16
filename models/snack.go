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
	Topics    []Topic  `json:"topics" db:"topics" many_to_many:"topics_snacks"`
	Authors   []Author `json:"authors" db:"authors" many_to_many:"authors_snacks"`
}

func init() {
	registry["snack"] = func() (IDer, interface{}) {
		return new(Snack), new([]Snack)
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
