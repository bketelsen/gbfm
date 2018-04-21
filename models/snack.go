package models

import "github.com/jinzhu/gorm"

// Snack is a snack - a small piece of news
type Snack struct {
	gorm.Model
	Slug string `json:"slug" db:"slug"`

	Title     string   `json:"title" db:"title"`
	Sponsored bool     `json:"sponsored" db:"sponsored"`
	URL       string   `json:"url" db:"url"`
	Summary   string   `json:"summary" db:"summary"`
	Comment   string   `json:"comment" db:"comment"`
	EmbedCode string   `json:"embed_code" db:"embed_code"`
	Topics    []Topic  `gorm:"many2many:topics_snacks;"`
	Authors   []Author `gorm:"many2many:authors_snacks;"`
}

/*
func init() {
	registry["snack"] = &registryFuncs{
		list:  func() Lister { return new(Snacks) },
		empty: func() IDer { return new(Snack) },
		sample: func() IDer {
			return &Snack{
				Slug:      namer.NameSep("-"),
				Title:     namer.Name(),
				Sponsored: true,
				URL:       namer.NameSep("-"),
				Summary:   namer.Name(),
				Comment:   namer.Name(),
				EmbedCode: namer.NameSep("-"),
			}
		},
	}
}

func (a Snack) GetID() uuid.UUID {
	return a.ID
}

func (a Snack) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Snack) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Snack) GetSlug() string {
	return a.Slug
}

// ModelName implements ModelNamer
func (Snack) ModelName() string {
	return "Snack"
}

// Snacks is a list of Snack models. It implements Lister
type Snacks []*Snack

// Len implements Lister
func (s Snacks) Len() int {
	return len(s)
}

// EltAt implements Lister
func (s Snacks) EltAt(i int) IDer {
	if i < len(s) {
		return s[i]
	}
	return nil
}
*/
