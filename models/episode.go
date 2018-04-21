package models

import "github.com/jinzhu/gorm"

// Episode represents an episode
type Episode struct {
	gorm.Model
	Slug string `json:"slug" db:"slug"`

	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Markdown    string `json:"markdown" db:"markdown"`

	Picture   ContentImageStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	EmbedCode string              `json:"embed_code" db:"embed_code"`
	Body      string              `json:"body" db:"body"`
	Pro       bool                `json:"pro" db:"pro"`
	Repo      string              `json:"repo" db:"repo"`
	Topics    []Topic             `gorm:"many2many:topics_episodes;"`
	Authors   []Author            `gorm:"many2many:authors_episodes;"`
	Series    []Series            `gorm:"many2many:episodes_series;"`
}

// BeforeSave is called before record saves.
// Sets the Slug from the Episode's Title
func (episode *Episode) BeforeSave(scope *gorm.Scope) (err error) {
	scope.SetColumn("Slug", sluggify(episode.Title))
	return nil
}
