package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// GIFM is a go in 5 minutes entry
type GIFM struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Slug        string    `json:"slug" db:"slug"`
	Title       string    `json:"title" db:"title"`
	EmdedCode   string    `json:"emded_code" db:"embed_code"`
	GithubLink  string    `json:"github_link" db:"github_link"`
	Sponsor     string    `json:"sponsor" db:"sponsor"`
	Description string    `json:"description" db:"description"`
	Topics      []Topic   `json:"topics" db:"topics" many_to_many:"gifm_topics"`
	Authors     []Author  `json:"authors" db:"authors" many_to_many:"gifm_authors"`
}

func init() {
	registry["gifm"] = func() (IDer, interface{}) {
		return new(GIFM), new([]GIFM)
	}
}

func (a GIFM) GetID() uuid.UUID {
	return a.ID
}

func (a GIFM) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a GIFM) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a GIFM) GetSlug() string {
	return a.Slug
}
