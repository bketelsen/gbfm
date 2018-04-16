package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// Author is an author
type Author struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Slug      string    `json:"slug" db:"slug"`

	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo_url" db:"photo_url"`
	// TODO: has_many's for the content models
}

// GetAuthor returns an author by the given ID
func GetAuthor(tx *pop.Connection, id uuid.UUID) (*Author, error) {
	a := new(Author)
	if err := tx.Eager().Where("id = ?", id).First(a); err != nil {
		return nil, err
	}
	return a, nil
}

func init() {
	registry["author"] = func() (IDer, interface{}) {
		return new(Author), new([]Author)
	}
}
func (a Author) GetID() uuid.UUID {
	return a.ID
}

func (a Author) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Author) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Author) GetSlug() string {
	return a.Slug
}
