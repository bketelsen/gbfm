package models

import (
	"github.com/blevesearch/bleve/mapping"
)

// implements https://godoc.org/github.com/ponzu-cms/ponzu/system/item#Sluggable
type ponzuSluggable struct {
	Slug string `json:"slug" db:"slug"`
}

func (p *ponzuSluggable) SetSlug(s string) {
	p.Slug = s
}

func (p *ponzuSluggable) ItemSlug() string {
	return p.Slug
}

// implements ItemID and SetItemID in
// https://godoc.org/github.com/ponzu-cms/ponzu/system/item#Identifiable
type ponzuPartialIdentifiable struct {
	ID int
}

func (p *ponzuPartialIdentifiable) ItemID() int {
	return p.ID
}

func (p *ponzuPartialIdentifiable) SetItemID(i int) {
	p.ID = i
}

type ponzuSearchable struct{}

func (p *ponzuSearchable) SearchMapping() (*mapping.IndexMappingImpl, error) {
	return nil, nil
}

func (p *ponzuSearchable) IndexContent() bool {
	return false
}
