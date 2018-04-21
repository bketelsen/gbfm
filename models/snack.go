package models

import "github.com/jinzhu/gorm"

// Snack is a snack - a small piece of news
type Snack struct {
	gorm.Model
	Slug      string      `json:"slug" db:"slug"`
	Title     string      `json:"title" db:"title"`
	Sponsored bool        `json:"sponsored" db:"sponsored"`
	URL       string      `json:"url" db:"url"`
	Summary   string      `json:"summary" db:"summary"`
	Comment   string      `json:"comment" db:"comment"`
	Icon      IconStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	EmbedCode string      `json:"embed_code" db:"embed_code"`
	Topics    []Topic     `gorm:"many2many:topics_snacks;"`
	Authors   []Author    `gorm:"many2many:authors_snacks;"`
}
