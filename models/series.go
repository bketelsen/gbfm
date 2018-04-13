package models

type Series struct {
	coreModel
	slugger
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Body        string   `json:"body"`
	Pro         bool     `json:"pro"`
	Topics      []Topic  `json:"topics" db:"topics" many_to_many:"topics_series"`
	Authors     []Author `json:"authors" db:"authors" many_to_many:"authors_series"`
}

func init() {
	registry["series"] = func() (interface{}, interface{}) {
		return new(Series), new([]Series)
	}
}
