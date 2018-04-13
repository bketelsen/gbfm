package models

// GIFM is a go in 5 minutes entry
type GIFM struct {
	coreModel
	slugger
	Title       string   `json:"title" db:"title"`
	EmdedCode   string   `json:"emded_code" db:"embed_code"`
	GithubLink  string   `json:"github_link" db:"github_link"`
	Sponsor     string   `json:"sponsor" db:"sponsor"`
	Description string   `json:"description" db:"description"`
	Topics      []Topic  `json:"topics" db:"topics" many_to_many:"gifm_topics"`
	Authors     []Author `json:"authors" db:"authors" many_to_many:"gifm_authors"`
}

func init() {
	registry["gifm"] = func() (interface{}, interface{}) {
		return new(GIFM), new([]GIFM)
	}
}
