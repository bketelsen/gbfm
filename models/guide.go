package models

import "github.com/jinzhu/gorm"

// Guide is a guide
type Guide struct {
	gorm.Model
	Slug         string   `json:"slug" db:"slug"`
	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description" db:"description"`
	Markdown     string   `json:"markdown" db:"markdown"`
	ThumbnailURL string   `json:"thumbnail_url" db:"thumbnail_url"`
	EmbedCode    string   `json:"embed_code" db:"embed_code"`
	Body         string   `json:"body" db:"body"`
	Pro          bool     `json:"pro" db:"pro"`
	Topics       []Topic  `gorm:"many2many:topics_guides;"`
	Authors      []Author `gorm:"many2many:authors_guides;"`
}

/*
func init() {
	registry["guide"] = &registryFuncs{
		empty: func() IDer { return new(Guide) },
		sample: func() IDer {
			return &Guide{
				Slug:         namer.NameSep("-"),
				Title:        namer.Name(),
				Description:  namer.Name(),
				Markdown:     namer.Name(),
				ThumbnailURL: namer.Name(),
				EmbedCode:    namer.Name(),
				Body:         namer.Name(),
				Pro:          true,
			}
		},
		list: func() Lister { return new(Guides) },
	}
}

func (a Guide) GetID() uuid.UUID {
	return a.ID
}

func (a Guide) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Guide) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Guide) GetSlug() string {
	return a.Slug
}

// ModelName implements ModelNamer
func (Guide) ModelName() string {
	return "Guide"
}

// Guides is a list of Guide models. It implements Lister
type Guides []Guide

// Len implements Lister
func (g Guides) Len() int {
	return len(g)
}

// EltAt implements Lister
func (g Guides) EltAt(i int) IDer {
	if i < len(g) {
		return g[i]
	}
	return nil
}
*/
