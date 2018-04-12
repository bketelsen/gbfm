package models

// Topic is a DB model for a topic
type Topic struct {
	coreModel
	slugger
	Name string `json:"name" db:"name"`
}
