package models

// Topic is a DB model for a topic
type Topic struct {
	coreModel
	slugger
	Name string `json:"name" db:"name"`
}

func init() {
	registry["topic"] = func() (interface{}, interface{}) {
		return new(Topic), new([]Topic)
	}
}
