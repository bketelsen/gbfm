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

// BeforeSave is called before record saves.
// Sets the Slug from the Series's Title
func (s *Series) BeforeSave(scope *gorm.Scope) (err error) {
	scope.SetColumn("Slug", sluggify(s.Title))
	return nil
}
