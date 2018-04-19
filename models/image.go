package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Image struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Slug        string       `json:"slug" db:"slug"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	AltText     string       `json:"alt_text" db:"alt_text"`
	File        binding.File `json:"file" db:"file"`
}

// String is not required by pop and may be deleted
func (i Image) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Images is not required by pop and may be deleted
type Images []Image

// String is not required by pop and may be deleted
func (i Images) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Image) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Title, Name: "Title"},
		&validators.StringIsPresent{Field: i.Description, Name: "Description"},
		&validators.StringIsPresent{Field: i.AltText, Name: "AltText"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Image) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Image) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
func init() {
	registry["image"] = &registryFuncs{
		empty: func() IDer { return new(Image) },
		list:  func() interface{} { return new([]Image) },
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
