package models

// Guide is a guide
type Guide struct {
	coreModel
	slugger
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	EmbedCode   string   `json:"embed_code"`
	Body        string   `json:"body"`
	Pro         bool     `json:"pro"`
	Topics      []Topic  `json:"topics" db:"topics" many_to_many:"topics_guides"`
	Authors     []Author `json:"authors" db:"authors" many_to_many:"authors_guides"`
}

func init() {
	registry["guide"] = func() (interface{}, interface{}) {
		return new(Guide), new([]Guide)
	}
}
