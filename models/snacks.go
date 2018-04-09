package models

import (
	"github.com/gophersnacks/content/content"
)

// FullSnack is a complete snack and its authors
type FullSnack struct {
	Snack      *content.Snack
	AuthorList []*content.Author
}

// GetFullSnackBySlug gets a full snack, including author list, by snack slug
func GetFullSnackBySlug(id string) (*FullSnack, error) {
	var fsn FullSnack
	sn, err := GetSnackBySlug(id)
	if err != nil {
		return nil, err
	}
	fsn = FullSnack{
		Snack:      &sn,
		AuthorList: make([]*content.Author, len(sn.Authors)),
	}
	for i, id := range AuthorIDsForSnack(sn) {
		a, err := GetAuthor(id)
		if err != nil {
			return nil, err
		}
		fsn.AuthorList[i] = &a
	}
	return &fsn, nil
}

// AuthorIDsForSnack returns the IDs of all of a Snack's authors
func AuthorIDsForSnack(s content.Snack) []int {
	var authors []int
	for _, s := range s.Authors {
		i, err := getID(s)
		if err == nil {
			authors = append(authors, i)
		}
	}
	return authors
}
