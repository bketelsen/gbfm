package models

// Snack is a snack - a small piece of news
type Snack struct {
	coreModel
	slugger

	Title     string   `json:"title"`
	Sponsored bool     `json:"sponsored"`
	URL       string   `json:"url"`
	Summary   string   `json:"summary"`
	Comment   string   `json:"comment"`
	SnackSlug string   `json:"snack_slug"`
	Topics    []Topic  `json:"topics" db:"topics" has_many:"topics"`
	Authors   []Author `json:"authors" db:"authors" has_many:"authors"`
}
