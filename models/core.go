package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// IDer is a model that exposes its ID
type IDer interface {
	GetID() uuid.UUID
}

// StringIDer is a model that exposes its ID and a string representation
type StringIDer interface {
	GetID() uuid.UUID
	String() string
}

// ModelNamer is a model that returns its name, without reflection.
// It's used in the formID template helper in actions/content/app.go
type ModelNamer interface {
	ModelName() string
}

// Lister is a list of a given model that can report on its length
type Lister interface {
	Len() int
	// EltAt returns the IDer at index i, or nil
	EltAt(i int) StringIDer
}

// Core represents the core of a model
type Core interface {
	IDer
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
type Topicer interface {
	GetTopics() []uuid.UUID
	AddTopics([]uuid.UUID, *pop.Connection) error
}

// implementation of Core
type coreModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (c coreModel) GetID() uuid.UUID {
	return c.ID
}

func (c coreModel) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c coreModel) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

// Slugger is a model that has a slug
type Slugger interface {
	GetSlug() string
}

// implementation of Slugger
type slugger struct {
	Slug string `json:"slug" db:"slug"`
}

func (s slugger) GetSlug() string {
	return s.Slug
}
