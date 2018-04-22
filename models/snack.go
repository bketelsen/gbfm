package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// Snack is a snack - a small piece of news
type Snack struct {
	gorm.Model
	Slug      string `json:"slug" db:"slug"`
	Title     string `json:"title" db:"title"`
	Source    string
	Sponsored bool `json:"sponsored" db:"sponsored"`
	Featured  bool

	Picture   ContentImageStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	URL       string              `json:"url" db:"url"`
	Summary   string              `json:"summary" db:"summary"`
	Comment   string              `json:"comment" db:"comment"`
	EmbedCode string              `json:"embed_code" db:"embed_code"`
	Topics    []Topic             `gorm:"many2many:topics_snacks;"`
	Authors   []Author            `gorm:"many2many:authors_snacks;"`
}

func (s Snack) DisplayAuthors() string {
	fmt.Println("DisplayAuthors")
	var authors []string
	if len(s.Authors) > 0 {
		for _, a := range s.Authors {
			authors = append(authors, a.Name)
		}
		fmt.Println("authors", authors)
		return strings.Join(authors, ", ")
	}
	return "GopherSnacks Staff"
}

// BeforeSave is called before record saves.
// Sets the Slug from the Snacks's Title
func (s *Snack) BeforeSave(scope *gorm.Scope) (err error) {
	scope.SetColumn("Slug", sluggify(s.Title))
	return nil
}
