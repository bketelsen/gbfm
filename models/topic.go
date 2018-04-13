package models

// Topic is a DB model for a topic
type Topic struct {
	coreModel
	slugger
	Name string `json:"name" db:"name"`
	// TODO: has_many's for the content models
}

func init() {
	registry["topic"] = func() (interface{}, interface{}) {
		return new(Topic), new([]Topic)
	}
}
