package models

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// Author is an author
type Author struct {
	coreModel
	slugger
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

// GetAuthor returns an author by the given ID
func GetAuthor(tx *pop.Connection, id uuid.UUID) (*Author, error) {
	a := new(Author)
	if err := tx.Eager().Where("id = ?", id).First(a); err != nil {
		return nil, err
	}
	return a, nil
}
