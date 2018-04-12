package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Core represents the core of a model
type Core interface {
	GetID() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
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

// Templater lets you define the templates that the model should use
type Templater interface {
	GetShowTemplateName() string
	GetEditTemplateName() string
	GetNewTemplateName() string
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

type CoreTemplater interface {
	Core
	Templater
}
