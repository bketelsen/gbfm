package models

import "github.com/jinzhu/gorm"

type Series struct {
	gorm.Model
	Slug        string              `json:"slug" db:"slug"`
	Title       string              `json:"title" db:"title"`
	Description string              `json:"description" db:"description"`
	Picture     ContentImageStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	Body        string              `json:"body" db:"body"`
	Pro         bool                `json:"pro" db:"pro"`

	Topics   []Topic   `gorm:"many2many:topics_series;"`
	Authors  []Author  `gorm:"many2many:authors_series;"`
	Episodes []Episode `gorm:"many2many:episodes_series;"`
}
