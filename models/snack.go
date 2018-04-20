package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// Snack is a snack - a small piece of news
type Snack struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Slug      string    `json:"slug" db:"slug"`

	Title     string   `json:"title" db:"title"`
	Sponsored bool     `json:"sponsored" db:"sponsored"`
	URL       string   `json:"url" db:"url"`
	Summary   string   `json:"summary" db:"summary"`
	Comment   string   `json:"comment" db:"comment"`
	EmbedCode string   `json:"embed_code" db:"embed_code"`
	Topics    []Topic  `json:"topics" many_to_many:"topics_snacks"`
	Authors   []Author `json:"authors" many_to_many:"authors_snacks"`
}

func init() {
	registry["snack"] = &registryFuncs{
		list:  func() Lister { return new(Snacks) },
		empty: func() IDer { return new(Snack) },
		sample: func() IDer {
			return &Snack{
				Slug:      namer.NameSep("-"),
				Title:     namer.Name(),
				Sponsored: true,
				URL:       namer.NameSep("-"),
				Summary:   namer.Name(),
				Comment:   namer.Name(),
				EmbedCode: namer.NameSep("-"),
			}
		},
	}
}

func (s *Snack) String() string {
	return s.Title
}

func (a Snack) GetID() uuid.UUID {
	return a.ID
}

func (a Snack) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a Snack) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a Snack) GetSlug() string {
	return a.Slug
}

// ModelName implements ModelNamer
func (Snack) ModelName() string {
	return "Snack"
}

// Snacks is a list of Snack models. It implements Lister
type Snacks []*Snack

// Len implements Lister
func (s Snacks) Len() int {
	return len(s)
}

// EltAt implements Lister
func (s Snacks) EltAt(i int) StringIDer {
	if i < len(s) {
		return s[i]
	}
	return nil
}

func (s *Snack) GetTopics() []uuid.UUID {
	var ids = make([]uuid.UUID, len(s.Topics))
	for x, t := range s.Topics {
		ids[x] = t.ID
	}
	return ids
}

func (s *Snack) AddTopics(ids []uuid.UUID, tx *pop.Connection) error {

	for _, id := range ids {
		topic := new(Topic)
		err := tx.Find(topic, id)
		if err != nil {
			return err
		}
		s.Topics = append(s.Topics, *topic)
	}
	return tx.Save(s)
}
