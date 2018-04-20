package models

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
)

type Image struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Slug        string       `json:"slug" db:"slug"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	AltText     string       `json:"alt_text" db:"alt_text"`
	File        binding.File `json:"-" db:"-" form:"File"`
	FileName    string       `json:"file_name" db:"file_name"`
}

// String is not required by pop and may be deleted
func (i Image) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// BeforeCreate sets the image's filename field
func (i *Image) BeforeCreate(tx *pop.Connection) error {
	if !i.File.Valid() {
		return nil
	}
	i.FileName = i.File.Filename
	return nil
}

// AfterCreate saves the file
func (i *Image) AfterCreate(tx *pop.Connection) error {
	if !i.File.Valid() {
		return nil
	}
	dir := filepath.Join(".", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, i.File.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, i.File)
	return err
}

// ModelName implements ModelNamer
func (Image) ModelName() string {
	return "Image"
}

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
func (i Image) GetID() uuid.UUID {
	return i.ID
}

// GetCreatedAt implements Core
func (i Image) GetCreatedAt() time.Time {
	return i.CreatedAt
}

// GetUpdatedAt implements Core
func (i Image) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}

// GetSlug implements Slugger
func (i Image) GetSlug() string {
	return i.Slug
}
