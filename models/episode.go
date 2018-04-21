package models

import "github.com/jinzhu/gorm"

// Episode represents an episode
type Episode struct {
	gorm.Model
	Slug string `json:"slug" db:"slug"`

	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description" db:"description"`
	Markdown     string   `json:"markdown" db:"markdown"`
	ThumbnailURL string   `json:"thumbnail_url" db:"thumbnail_url"`
	EmbedCode    string   `json:"embed_code" db:"embed_code"`
	Body         string   `json:"body" db:"body"`
	Pro          bool     `json:"pro" db:"pro"`
	Repo         string   `json:"repo" db:"repo"`
	Topics       []Topic  `gorm:"many2many:topics_episodes;"`
	Authors      []Author `gorm:"many2many:authors_episodes;"`
	Series       []Series `gorm:"many2many:episodes_series;"`
}

/*
func init() {
	registry["episode"] = &registryFuncs{
		empty: func() IDer { return new(Episode) },
		list:  func() Lister { return new(Episodes) },
		sample: func() IDer {
			return &Episode{
				Slug:         namer.NameSep("-"),
				Title:        namer.Name(),
				Description:  namer.Name(),
				ThumbnailURL: namer.NameSep("-"),
				EmbedCode:    namer.NameSep("-"),
				Body:         namer.Name(),
				Pro:          true,
				Repo:         namer.NameSep("-"),
			}
		},
	}
}

// GetID implements IDer
func (a Episode) GetID() uuid.UUID {
	return a.ID
}

// GetCreatedAt implements Core
func (a Episode) GetCreatedAt() time.Time {
	return a.CreatedAt
}

// GetUpdatedAt implements Core
func (a Episode) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

// GetSlug implements Slugger
func (a Episode) GetSlug() string {
	return a.Slug
}

// ModelName implements ModelNamer
func (Episode) ModelName() string {
	return "Episode"
}

// Episodes is a list of Authors. It implements Lister
type Episodes []*Episode

// Len implements Lister
func (e Episodes) Len() int {
	return len(e)
}

// EltAt implements Lister
func (e Episodes) EltAt(i int) IDer {
	if i < len(e) {
		return e[i]
	}
	return nil
}
*/
