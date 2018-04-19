package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/uuid"
)

type Image struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Slug        string       `json:"slug" db:"slug"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	AltText     string       `json:"alt_text" db:"alt_text"`
	File        binding.File `json:"file" db:"-" form:"File"`
}

// String is not required by pop and may be deleted
func (i Image) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Images is not required by pop and may be deleted
type Images []Image

// ImageList is a list of Image models. It implements Lister
type ImageList []*Image

// Len implements Lister
func (s ImageList) Len() int {
	return len(s)
}

// EltAt implements Lister
func (s ImageList) EltAt(i int) IDer {
	if i < len(s) {
		return s[i]
	}
	return nil
}

// String is not required by pop and may be deleted
func (i Images) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

func init() {
	registry["image"] = &registryFuncs{
		empty: func() IDer { return new(Image) },
		list:  func() Lister { return new(ImageList) },
		sample: func() IDer {
			return &Image{
				Slug:        namer.NameSep("-"),
				Title:       namer.Name(),
				Description: namer.Name(),
				AltText:     namer.Name(),
			}
		},
	}
}

// GetID implements IDer
func (a Image) GetID() uuid.UUID {
	return a.ID
}

// GetCreatedAt implements Core
func (a Image) GetCreatedAt() time.Time {
	return a.CreatedAt
}

// GetUpdatedAt implements Core
func (a Image) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

// GetSlug implements Slugger
func (a Image) GetSlug() string {
	return a.Slug
}
