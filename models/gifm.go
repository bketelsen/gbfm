package models

// GIFM is a go in 5 minutes entry
type GIFM struct {
	coreModel
	slugger
	Title       string   `json:"title"`
	EmdedCode   string   `json:"emded_code"`
	GithubLink  string   `json:"github_link"`
	Sponsor     string   `json:"sponsor"`
	Description string   `json:"description"`
	Topics      []Topic  `json:"topics" db:"topics" many_to_many:"topics_gifm"`
	Authors     []Author `json:"authors" db:"authors" many_to_many:"authors_gifm"`
}

func init() {
	registry["gifm"] = func() (interface{}, interface{}) {
		return new(GIFM), new([]GIFM)
	}
}
