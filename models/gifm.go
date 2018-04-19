package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Gifm is a go in 5 minutes entry.
//
// This does not match the migrations. It's called Gifm and the migration
// creates a "gbfm" table. TODOs:
//
// - Rename this GBFM
// - Change the TableName func to return "gbfm"
//
type Gifm struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Slug        string    `json:"slug" db:"slug"`
	Title       string    `json:"title" db:"title"`
	EmdedCode   string    `json:"emded_code" db:"embed_code"`
	GithubLink  string    `json:"github_link" db:"github_link"`
	Sponsor     string    `json:"sponsor" db:"sponsor"`
	Description string    `json:"description" db:"description"`
	Topics      []Topic   `json:"topics" many_to_many:"gifm_topics"`
	Authors     []Author  `json:"authors" many_to_many:"gifm_authors"`
}

func init() {
	registry["gifm"] = &registryFuncs{
		list:  func() Lister { return new(Gifms) },
		empty: func() IDer { return new(Gifm) },
		sample: func() IDer {
			return &Gifm{
				Slug:        namer.NameSep("-"),
				Title:       namer.Name(),
				EmdedCode:   namer.NameSep("-"),
				GithubLink:  namer.NameSep("-"),
				Sponsor:     namer.Name(),
				Description: namer.Name(),
			}
		},
	}
}

// TableName implements the pop TableNamer interface. This needs to be
// implemented because pop will automatically infer the table name from
// the struct name as "g_i_f_ms". This is because it separates capital
// letters with underscores
func (a Gifm) TableName() string {
	return "gbfms"
}
func (a Gifm) GetID() uuid.UUID {
	return a.ID
}

func (a Gifm) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Gifm) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Gifm) GetSlug() string {
	return a.Slug
}

// Gifms is a list of GIFM models. It implements Lister
type Gifms []*Gifm

// Len implements Lister
func (g Gifms) Len() int {
	return len(g)
}

// EltAt implements Lister
func (g Gifms) EltAt(i int) IDer {
	if i < len(g) {
		return g[i]
	}
	return nil
}
