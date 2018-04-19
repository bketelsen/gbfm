package models

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// Gbfm is a go in 5 minutes entry.
type Gbfm struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Slug        string    `json:"slug" db:"slug"`
	Title       string    `json:"title" db:"title"`
	EmdedCode   string    `json:"emded_code" db:"embed_code"`
	GithubLink  string    `json:"github_link" db:"github_link"`
	Sponsor     string    `json:"sponsor" db:"sponsor"`
	Description string    `json:"description" db:"description"`
	Topics      []Topic   `json:"topics" many_to_many:"gbfm_topics"`
	Authors     []Author  `json:"authors" many_to_many:"gbfm_authors"`
}

func init() {
	registry["gbfm"] = &registryFuncs{
		list:  func() Lister { return new(Gbfms) },
		empty: func() IDer { return new(Gbfm) },
		sample: func() IDer {
			return &Gbfm{
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
func (a Gbfm) TableName() string {
	return "gbfms"
}
func (a Gbfm) GetID() uuid.UUID {
	return a.ID
}

func (a Gbfm) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Gbfm) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Gbfm) GetSlug() string {
	return a.Slug
}

// Gbfms is a list of GBFM models. It implements Lister
type Gbfms []*Gbfm

// Len implements Lister
func (g Gbfms) Len() int {
	return len(g)
}

// EltAt implements Lister
func (g Gbfms) EltAt(i int) IDer {
	if i < len(g) {
		return g[i]
	}
	return nil
}
