package models

import "github.com/jinzhu/gorm"

// Guide is a guide
type Guide struct {
	gorm.Model
	Slug        string              `json:"slug" db:"slug"`
	Title       string              `json:"title" db:"title"`
	Description string              `json:"description" db:"description"`
	Markdown    string              `json:"markdown" db:"markdown"`
	Picture     ContentImageStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	EmbedCode   string              `json:"embed_code" db:"embed_code"`
	Body        string              `json:"body" db:"body"`
	Pro         bool                `json:"pro" db:"pro"`
	Topics      []Topic             `gorm:"many2many:topics_guides;"`
	Authors     []Author            `gorm:"many2many:authors_guides;"`
}
