package models

import (
	"github.com/jinzhu/gorm"
)

// Author is an author
type Author struct {
	gorm.Model
	Slug string `json:"slug" db:"slug"`

	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	Avatar AvatarImageStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	User   User
	UserID uint
	// TODO: has_many's for the content models

}

// BeforeSave is called before record saves.
// Sets the Slug from the Author's Name
func (author *Author) BeforeSave(scope *gorm.Scope) (err error) {
	scope.SetColumn("Slug", sluggify(author.Name))
	return nil
}
