package models

// Snack is a snack - a small piece of news
type Snack struct {
	coreModel
	slugger

	Title     string   `json:"title" db:"title"`
	Sponsored bool     `json:"sponsored" db:"sponsored"`
	URL       string   `json:"url" db:"url"`
	Summary   string   `json:"summary" db:"summary"`
	Comment   string   `json:"comment" db:"summary"`
	Topics    []Topic  `json:"topics" db:"topics" many_to_many:"topics_snacks"`
	Authors   []Author `json:"authors" db:"authors" many_to_many:"authors_snacks"`
}

func init() {
	registry["snack"] = func() (interface{}, interface{}) {
		return new(Snack), new([]Snack)
	}
}
