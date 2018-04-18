package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Topic is a DB model for a topic
type Topic struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Slug      string    `json:"slug" db:"slug"`
	Name      string    `json:"name" db:"name"`
	// TODO: has_many's for the content models
}

func init() {
	registry["topic"] = &registryFuncs{
		empty: func() IDer { return new(Topic) },
		list:  func() interface{} { return new([]Topic) },
		sample: func() IDer {
			return &Topic{
				Slug: namer.NameSep("-"),
				Name: namer.Name(),
			}
		},
	}
}

func (a Topic) GetID() uuid.UUID {
	return a.ID
}

func (a Topic) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Topic) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Topic) GetSlug() string {
	return a.Slug
}
