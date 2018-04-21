package models

import "github.com/jinzhu/gorm"

// Topic is a DB model for a topic
type Topic struct {
	gorm.Model
	Slug     string      `json:"slug" db:"slug"`
	Name     string      `json:"name" db:"name"`
	Icon     IconStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	Snacks   []Snack     `gorm:"many2many:topics_snacks;"`
	Episodes []Episode   `gorm:"many2many:topics_episodes;"`
	Guides   []Guide     `gorm:"many2many:topics_guides;"`
}
